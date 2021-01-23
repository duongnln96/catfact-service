package factfinder

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// ServiceInfo struct contain information about service
type ServiceInfo struct {
	ServiceID string `mapstructure:"service"`
	Version   string `mapstructure:"version"`
}

// CoreConfig struct
type CoreConfig struct {
	LocalPort     int    `mapstructure:"local-port"`
	LocalProtocal string `mapstructure:"local-protocol"`
	OfflineMode   bool   `mapstructure:"offline-mode"`
}

type jsonInterface struct {
	Text string `json:"text"`
}

type catFact struct {
	Text string `json:"text"`
}

// CoreFactFinderConfig include all config for processor
type CoreFactFinderConfig struct {
	Log           *zap.SugaredLogger
	ModeOffline   bool
	Port          int
	LocalProtocal string
}

// CoreProcessor class
type CoreProcessor struct {
	log           *zap.SugaredLogger
	mode          bool
	port          int
	localProtocal string
}

// ICoreFactFinder interface
type ICoreFactFinder interface {
	Start()
	Stop()
}

// NewCoreFactFinder function
func NewCoreFactFinder(cfg CoreFactFinderConfig) ICoreFactFinder {
	return &CoreProcessor{
		log:           cfg.Log,
		mode:          cfg.ModeOffline,
		port:          cfg.Port,
		localProtocal: cfg.LocalProtocal,
	}
}

// Start method
func (core *CoreProcessor) Start() {
	core.log.Info("Start factfinder service")

	http.HandleFunc("/health", healthCheck)

	localPort := fmt.Sprintf(":%d", core.port)

	core.log.Fatal(http.ListenAndServe(localPort, nil))

}

// Stop method
func (core *CoreProcessor) Stop() {
	core.log.Info("Stop factfinder service")
}
