package cmd

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"github.com/jitsucom/jitsu/server/telemetry"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	maxChunkSize = 20 * 1024 * 1024 // 20 MB
	dateLayout   = "2006-01-02"
)

var (
	//command flags
	state, start, end, host, apiKey string
	chunkSize                       int64
	disableProgressBars             string
	//command args
	files []string
)

// replayCmd represents the base command when called without any subcommands
var replayCmd = &cobra.Command{
	Use:   "replay [flags] <files>",
	Short: "CLI for uploading data from local files into Jitsu destinations via API",
	Long:  `Jitsu CLI tool for bulk uploading files with events into Jitsu. Common use case: upload archive logs (aka replay)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("SERVER_TELEMETRY_DISABLED_USAGE") != "true" {
			var cs int64
			if chunkSize != maxChunkSize {
				cs = chunkSize
			}
			telemetry.CLIStart("replay", start != "" || end != "", state != "", cs)
			telemetry.Flush()
			time.Sleep(time.Second)
		}

		if len(args) == 0 {
			return errors.New("requires at least 1 file as an arg")
		}
		return replay(args)
	},
	Version: version,
}

func init() {
	rootCmd.AddCommand(replayCmd)

	replayCmd.Flags().StringVar(&state, "state", "", "(optional) a path to file where Jitsu will save the state - already uploaded files names. It prevents resending already loaded files on each run")
	replayCmd.Flags().StringVar(&start, "start", "", "(optional) start date as YYYY-MM-DD. Treated as the beginning of the day UTC (YYYY-MM-DD 00:00:00.000Z). If missing, all files will be processed")
	replayCmd.Flags().StringVar(&end, "end", "", "(optional) end date as YYYY-MM-DD. Treated as the end of the day UTC (YYYY-MM-DD 23:59:59.999Z). If missing, all will be processed")
	replayCmd.Flags().StringVar(&host, "host", "http://localhost:8000", "(optional) Jitsu host")
	replayCmd.Flags().Int64Var(&chunkSize, "chunk-size", maxChunkSize, "(optional) max data chunk size in bytes (default 20 MB). If file size is greater then the file will be split into N chunks with max size and sent to Jitsu")
	replayCmd.Flags().StringVar(&disableProgressBars, "disable-progress-bars", "false", "(optional) if true then progress bars won't be displayed")

	replayCmd.Flags().StringVar(&apiKey, "api-key", "", "(required) Jitsu API Server secret. Data will be loaded into all destinations linked to this API Key.")
	replayCmd.MarkFlagRequired("api-key")
}

//replay is a command main function:
//reads files from filesystem and sends them to Jitsu
//operating:
// 1. always with full path filenames
// 2. always sends gzipped payloads to Jitsu
//returns err if occurred
func replay(inputFiles []string) error {
	matchedFiles, err := findFiles(inputFiles)
	if err != nil {
		return fmt.Errorf("find files error: %v", err)
	}
	absoluteFileNames, err := reformatFileNames(matchedFiles)
	if err != nil {
		return fmt.Errorf("preprocessing files failed: %v", err)
	}

	absoluteFileNamesAfterFiltering, err := filterFiles(absoluteFileNames, start, end)
	if err != nil {
		if err != nil {
			return fmt.Errorf("filtering files by date failed: %v", err)
		}
	}

	if len(absoluteFileNamesAfterFiltering) == 0 {
		return errors.New("none of the files match the --start --end condition")
	}

	if state != "" {
		var err error
		state, err = filepath.Abs(state)
		if err != nil {
			return fmt.Errorf("failed to get absolute state file path: %v", err)
		}
	}

	stateManager, err := newStateManager(state)
	if err != nil {
		return fmt.Errorf("error creating file state manager: %v", err)
	}

	var filesToUpload []string
	for _, f := range absoluteFileNamesAfterFiltering {
		//filter state file
		if f == state {
			continue
		}

		if !stateManager.IsUploaded(f) {
			filesToUpload = append(filesToUpload, f)
		}
	}

	if len(filesToUpload) == 0 {
		return errors.New("all files are marked as uploaded in state. Nothing to replay.")
	}

	var globalBar ProgressBar
	capacity := int64(len(filesToUpload))
	if disableProgressBars == "true" {
		globalBar = &DummyProgressBar{}
	} else {
		globalBar = NewParentMultiProgressBar(capacity)
	}

	client := newBulkClient(host, apiKey)

	var processedFiles int64
	for _, absFilePath := range filesToUpload {
		fileStat, err := os.Stat(absFilePath)
		if err != nil {
			return err
		}

		if err := uploadFile(globalBar, client, absFilePath, fileStat.Size()); err != nil {
			return fmt.Errorf("uploading file: %s\nmessage: %v", absFilePath, err)
		}
		processedFiles++
		globalBar.SetCurrent(processedFiles)
		stateManager.Success(absFilePath)
	}

	globalBar.SetCurrent(capacity)
	//wait for globalBar filled
	time.Sleep(time.Second)
	stateManager.Close()

	return nil
}

//uploadFile divides input file into chunks if size is grater then chunkSize
//sends data to Jitsu
//returns err if occurred
func uploadFile(globalBar ProgressBar, client *bulkClient, filePath string, fileSize int64) error {
	if fileSize > chunkSize {
		return sendChunked(globalBar, filePath, fileSize, client.sendGzippedMultiPart)
	}

	//send the whole file
	payload, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if filepath.Ext(filePath) != ".gz" {
		payload, err = doGzip(payload)
		if err != nil {
			return err
		}
	}

	//payload size of already gzipped
	payloadSize := int64(len(payload))
	processingTime := int64(float64(payloadSize) * 0.1)
	capacity := payloadSize + processingTime
	fileProgressBar := globalBar.createKBFileBar(filePath, capacity)
	if err := client.sendGzippedMultiPart(fileProgressBar, filePath, payload); err != nil {
		return err
	}

	fileProgressBar.SetCurrent(capacity)

	return nil
}

//sendChunked reads file maxChunkSize bytes and sends each chunk separately
func sendChunked(progressBar ProgressBar, filePath string, fileSize int64, sender func(fileProgressBar ProgressBar, filePath string, payload []byte) error) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	capacity := fileSize/chunkSize + 1
	var scanner *bufio.Scanner
	if filepath.Ext(filePath) == ".gz" {
		content, err := gzip.NewReader(file)
		if err != nil {
			return err
		}
		scanner = bufio.NewScanner(content)
	} else {
		scanner = bufio.NewScanner(file)
	}

	fileProgressBar := progressBar.createPartFileBar(filePath, capacity)

	cbuffer := make([]byte, 0, bufio.MaxScanTokenSize)
	scanner.Buffer(cbuffer, bufio.MaxScanTokenSize*100)

	chunk := bytes.Buffer{}
	var progress int64
	for scanner.Scan() {
		line := scanner.Bytes()
		if int64(chunk.Len()) > chunkSize {
			gzipped, err := doGzip(chunk.Bytes())
			if err != nil {
				return err
			}

			if err := sender(nil, filePath, gzipped); err != nil {
				return err
			}
			progress++
			fileProgressBar.SetCurrent(progress)
			chunk.Reset()
		}

		if chunk.Len() > 0 {
			chunk.Write([]byte("\n"))
		}

		if _, err := chunk.Write(line); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if chunk.Len() > 0 {
		//send
		gzipped, err := doGzip(chunk.Bytes())
		if err != nil {
			return err
		}

		if err := sender(nil, filePath, gzipped); err != nil {
			return err
		}
	}

	fileProgressBar.SetCurrent(capacity)

	return nil
}

//findFiles find files by masks and returns them
//if mask == filename then adds it to the result as well
func findFiles(masks []string) ([]string, error) {
	var fileNames []string
	for _, mask := range masks {
		matched, err := filepath.Glob(mask)
		if err != nil {
			return nil, err
		}

		fileNames = append(fileNames, matched...)
	}

	return fileNames, nil
}

//reformatFileNames returns files list with absolute path
//All directories in the list will be read recursively
func reformatFileNames(files []string) ([]string, error) {
	var result []string
	for _, file := range files {
		//skip mac os system files
		if strings.HasSuffix(file, ".DS_Store") {
			continue
		}

		f, err := filepath.Abs(file)
		if err != nil {
			return nil, fmt.Errorf("error getting absolute path for %s: %v", f, err)
		}

		if err := filepath.Walk(f,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				//skip mac os system files
				if strings.HasSuffix(path, ".DS_Store") {
					return nil
				}
				if !info.IsDir() {
					result = append(result, path)
				}

				return nil
			}); err != nil {
			return nil, err
		}
	}

	return result, nil
}

//filterFiles filters files by date and returns
func filterFiles(absoluteFileNames []string, startStr string, endStr string) ([]string, error) {
	if startStr == "" && endStr == "" {
		return absoluteFileNames, nil
	}

	startDate := time.Time{}
	if startStr != "" {
		t, err := time.Parse(dateLayout, startStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing 'start': %v", err)
		}
		startDate = t
	}

	endDate := time.Now().UTC()
	if endStr != "" {
		t, err := time.Parse(dateLayout, endStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing 'end': %v", err)
		}
		endDate = t.Add(time.Hour*23 + time.Minute*59 + time.Second*59 + time.Millisecond*999)
	}

	var result []string
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	for _, fn := range absoluteFileNames {
		filename := filepath.Base(fn)
		if re.MatchString(filename) {
			submatchall := re.FindAllString(filename, -1)
			for _, submatch := range submatchall {
				fileTime, err := time.Parse(dateLayout, submatch)
				if err != nil {
					return nil, fmt.Errorf("error parsing filename's [%s] date: %v", filename, err)
				}
				if startDate.Before(fileTime) && endDate.After(fileTime) {
					result = append(result, fn)
					break
				}
			}
		} else {
			fmt.Println(fmt.Sprintf("file %s doesn't contain date in its name. The file will be ignored.", fn))
		}
	}

	return result, nil
}

//doGzip returns gzipped payload
func doGzip(payload []byte) ([]byte, error) {
	gzipped := bytes.Buffer{}
	gzw := gzip.NewWriter(&gzipped)
	if _, err := io.Copy(gzw, bytes.NewBuffer(payload)); err != nil {
		return nil, err
	}

	if err := gzw.Close(); err != nil {
		return nil, err
	}

	return gzipped.Bytes(), nil
}
