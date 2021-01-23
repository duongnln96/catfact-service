package factfinder

import "go.uber.org/zap"

// Meta struct contain information about service
type Meta struct {
	serviceID string
	version   string
}

// CoreConfig class
type CoreConfig struct {
	modeOffline   bool
	port          int
	localProtocal string
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
