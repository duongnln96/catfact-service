package quotehandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var (
	serviceID string = "cat quote request"
	version   string = "v0.0.1"
)

// CoreCatfactQuoteConfig include the config from main app
type CoreCatfactQuoteConfig struct {
	Log                *zap.SugaredLogger
	Timeout            time.Duration
	LocalPort          int
	LocalProtocal      string
	FactFinderHost     string
	FactFinderPort     int
	FactFinderProtocol string
	FactFinderURI      string
}

// CoreProcessor class
type CoreProcessor struct {
	log                *zap.SugaredLogger
	timeout            time.Duration
	localPort          int
	localProtocal      string
	factFinderHost     string
	factFinderPort     int
	factFinderProtocol string
	factFinderURI      string
}

// ICoreCatfactQuote interface
type ICoreCatfactQuote interface {
	Start() error
	Stop()
}

// NewCoreCatQuote func
func NewCoreCatQuote(cfg CoreCatfactQuoteConfig) ICoreCatfactQuote {
	return &CoreProcessor{
		log:                cfg.Log,
		timeout:            cfg.Timeout,
		localPort:          cfg.LocalPort,
		localProtocal:      cfg.LocalProtocal,
		factFinderHost:     cfg.FactFinderHost,
		factFinderPort:     cfg.FactFinderPort,
		factFinderProtocol: cfg.FactFinderProtocol,
		factFinderURI:      cfg.FactFinderURI,
	}
}

type factFinderInterface struct {
	Text string `json:"text"`
}

type summary struct {
	EnglishPhrase string `json:"englishPhrase"`
}

func (c *CoreProcessor) apiHandler(resp http.ResponseWriter, req *http.Request) {
	quote, err := c.requestHandler()
	if err != nil {
		c.log.Error("Quote request processing failed")
	}
	resp.Header().Set("Content-Type", "text/html;charset=UTF-8")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(resp, "%v", quote)
	c.log.Infof("\nresq.body: %s \n Quote request successfully processed.", quote)
}

func (c *CoreProcessor) requestHandler() (string, error) {
	englishPhrase, err := c.findQuotes()
	if err != nil {
		return "", err
	}

	jsonString := &summary{
		EnglishPhrase: englishPhrase,
	}

	b, err := json.Marshal(&jsonString)
	if err != nil {
		c.log.Errorf("Fail to marshall text %+v", err)
		return "", err
	}

	return fmt.Sprintf("%s", string(b)), err
}

func (c *CoreProcessor) findQuotes() (string, error) {
	url := fmt.Sprintf("%s://%s:%d%s", c.factFinderProtocol, c.factFinderHost,
		c.factFinderPort, c.factFinderURI)
	c.log.Infof("url: %s", url)
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{
		Timeout: c.timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		c.log.Errorf("Fail to send HTTP request %+v", err)
		return "", err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var v factFinderInterface
	err = decoder.Decode(&v)
	if err != nil {
		c.log.Error("Failed to decode HTTP respond %+v", err)
		return "", err
	}

	return v.Text, err
}

func (c *CoreProcessor) healthCheck(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Ok")
}

// Start method
func (c *CoreProcessor) Start() error {
	c.log.Infof("Start %s:%s", serviceID, version)

	http.HandleFunc("/health", c.healthCheck)
	http.HandleFunc("/quote-request", c.apiHandler)

	localPort := fmt.Sprintf(":%d", c.localPort)

	err := http.ListenAndServe(localPort, nil)
	if err != nil {
		c.log.Fatalf("Start server fail: %+v", err)
		return err
	}

	return nil
}

// Stop method
func (c *CoreProcessor) Stop() {
	c.log.Info("Stop cat qoute service")
}
