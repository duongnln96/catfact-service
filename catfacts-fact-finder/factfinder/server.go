package factfinder

import "go.uber.org/zap"

// ServiceInfo struct contain information about service
type ServiceInfo struct {
	ServiceID string `mapstructure:"service"`
	Version   string `mapstructure:"version"`
}

// CoreConfig class
type CoreConfig struct {
	LocalPort     int    `mapstructure:"local-port"`
	LocalProtocal string `mapstructure:"local-protocol"`
	OfflineMode   bool   `mapstructure:"offline-mode"`
}

type coreFactFinderConfig struct {
	log           *zap.SugaredLogger
	modeOffline   bool
	port          int
	localProtocal string
}

// ICoreFactFinder interface
type ICoreFactFinder interface {
	Start()
	Stop()
}

func (ff *coreFactFinderConfig) Start() {
	ff.log.Info("Start factfinder service")
}

func (ff *coreFactFinderConfig) Stop() {
	ff.log.Info("Stop factfinder service")
}
