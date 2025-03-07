package adapters

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	amplitudeAPIURL = "https://api.amplitude.com/2/httpapi"
)

//AmplitudeRequest is a dto for sending requests to Amplitude
type AmplitudeRequest struct {
	APIKey string                   `json:"api_key"`
	Events []map[string]interface{} `json:"events"`
}

//AmplitudeResponse is a dto for receiving response from Amplitude
type AmplitudeResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

//AmplitudeRequestFactory is a factory for building Amplitude HTTP requests from input events
type AmplitudeRequestFactory struct {
	apiKey string
}

//newAmplitudeRequestFactory returns configured HTTPRequestFactory instance for amplitude requests
func newAmplitudeRequestFactory(apiKey string) (HTTPRequestFactory, error) {
	return &AmplitudeRequestFactory{apiKey: apiKey}, nil
}

//Create returns created amplitude request
//put empty array in body if object is nil (is used in test connection)
func (arf *AmplitudeRequestFactory) Create(object map[string]interface{}) (*Request, error) {
	//empty array is required. Otherwise nil will be sent (error)
	eventsArr := []map[string]interface{}{}
	if object != nil {
		eventsArr = append(eventsArr, object)
	}

	req := AmplitudeRequest{APIKey: arf.apiKey, Events: eventsArr}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling amplitude request [%v]: %v", req, err)
	}
	return &Request{
		URL:     amplitudeAPIURL,
		Method:  http.MethodPost,
		Body:    b,
		Headers: map[string]string{"Content-Type": "application/json"},
	}, nil
}

func (arf *AmplitudeRequestFactory) Close() {
}

//AmplitudeConfig is a dto for parsing Amplitude configuration
type AmplitudeConfig struct {
	APIKey string `mapstructure:"api_key" json:"api_key,omitempty" yaml:"api_key,omitempty"`
}

//Validate returns err if invalid
func (ac *AmplitudeConfig) Validate() error {
	if ac == nil {
		return errors.New("amplitude config is required")
	}
	if ac.APIKey == "" {
		return errors.New("'api_key' is required parameter")
	}

	return nil
}

//Amplitude is an adapter for sending HTTP requests to Amplitude
type Amplitude struct {
	AbstractHTTP

	config *AmplitudeConfig
}

//NewAmplitude returns configured Amplitude adapter instance
func NewAmplitude(config *AmplitudeConfig, httpAdapterConfiguration *HTTPAdapterConfiguration) (*Amplitude, error) {
	httpReqFactory, err := newAmplitudeRequestFactory(config.APIKey)
	if err != nil {
		return nil, err
	}

	httpAdapterConfiguration.HTTPReqFactory = httpReqFactory

	httpAdapter, err := NewHTTPAdapter(httpAdapterConfiguration)
	if err != nil {
		return nil, err
	}

	a := &Amplitude{config: config}
	a.httpAdapter = httpAdapter
	return a, nil
}

//NewTestAmplitude returns test instance of adapter
func NewTestAmplitude(config *AmplitudeConfig) *Amplitude {
	return &Amplitude{config: config}
}

//TestAccess sends test request (empty POST) to Amplitude and check if error has occurred
func (a *Amplitude) TestAccess() error {
	httpReqFactory, err := newAmplitudeRequestFactory(a.config.APIKey)
	if err != nil {
		return err
	}

	r, err := httpReqFactory.Create(nil)
	if err != nil {
		return err
	}

	httpReq, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return err
	}

	for k, v := range r.Headers {
		httpReq.Header.Add(k, v)
	}

	//send empty request and expect error
	resp, err := http.DefaultClient.Do(httpReq)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Error reading amplitude response body: %v", err)
		}

		response := &AmplitudeResponse{}
		err = json.Unmarshal(responseBody, response)
		if err != nil {
			return fmt.Errorf("Error unmarshalling amplitude response body: %v", err)
		}

		if response.Code != 200 {
			return fmt.Errorf("Error connecting to amplitude [code=%d]: %s", response.Code, response.Error)
		}

		//assume other errors - it's ok
		return nil
	}

	return err
}

//Type returns adapter type
func (a *Amplitude) Type() string {
	return "Amplitude"
}
