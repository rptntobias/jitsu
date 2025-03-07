package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jitsucom/jitsu/server/appconfig"
	"github.com/jitsucom/jitsu/server/logging"
	"github.com/jitsucom/jitsu/server/system"
	"github.com/spf13/viper"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	htmlContentType = "text/html; charset=utf-8"
	welcomePageName = "welcome.html"

	configuratorPresentKey = "__JITSU_CONFIGURATOR_PRESENT__"
	configuratorURLKey     = "__JITSU_CONFIGURATOR_URL__"
)

var blankPage = `<html><head><title>Jitsu edge server ver [VERSION]</title></head><body><pre><small><b>Jitsu edge server ver [VERSION]</b></small></pre></body></html>`

//RootPathHandler serves:
// HTTP redirect to Configurator
// HTML Welcome page or blanc page
type RootPathHandler struct {
	service         *system.Service
	configuratorURL string
	welcome         *template.Template
	redirectToHttps bool
}

//NewRootPathHandler reads sourceDir and returns RootPathHandler instance
func NewRootPathHandler(service *system.Service, sourceDir, configuratorURL string, disableWelcomePage, redirectToHttps bool) *RootPathHandler {
	rph := &RootPathHandler{service: service, configuratorURL: configuratorURL, redirectToHttps: redirectToHttps}

	if service.IsConfigured() {
		return rph
	}

	if disableWelcomePage {
		return rph
	}

	if !strings.HasSuffix(sourceDir, "/") {
		sourceDir += "/"
	}
	payload, err := ioutil.ReadFile(sourceDir + welcomePageName)
	if err != nil {
		logging.Errorf("Error reading %s file: %v", sourceDir+welcomePageName, err)
		return rph
	}

	welcomeHTMLTmpl, err := template.New("html template").
		Option("missingkey=zero").
		Parse(string(payload))
	if err != nil {
		logging.Error("Error parsing html template from", welcomePageName, err)
		return rph
	}

	rph.welcome = welcomeHTMLTmpl

	return rph
}

//Handler handles requests and returns welcome page or redirect to Configurator URL
func (rph *RootPathHandler) Handler(c *gin.Context) {
	if rph.service.ShouldBeRedirected() {
		redirectSchema := c.GetHeader("X-Forwarded-Proto")
		redirectHost := c.GetHeader("X-Forwarded-Host")
		realHost := c.GetHeader("X-Real-Host")
		if rph.redirectToHttps {
			//use X-Forwarded-Host if redirect to https
			//used in heroku deployments
			redirectSchema = "https"
			realHost = redirectHost
		}

		redirectURL := redirectSchema + "://" + realHost + viper.GetString("server.configurator_url")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
		return
	}

	c.Header("Content-type", htmlContentType)

	if rph.welcome == nil {
		c.Writer.Write([]byte(strings.ReplaceAll(blankPage,"[VERSION]", appconfig.RawVersion)))
		return
	}

	parameters := map[string]interface{}{configuratorPresentKey: false, configuratorURLKey: ""}
	if rph.configuratorURL != "" {
		parameters[configuratorURLKey] = rph.configuratorURL
		parameters[configuratorPresentKey] = true
	}

	err := rph.welcome.Execute(c.Writer, parameters)
	if err != nil {
		logging.Error("Error executing welcome.html template", err)
	}
}
