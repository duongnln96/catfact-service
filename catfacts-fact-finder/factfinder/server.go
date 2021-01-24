package factfinder

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	offlineMode   bool
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
		offlineMode:   cfg.ModeOffline,
		port:          cfg.Port,
		localProtocal: cfg.LocalProtocal,
	}
}

type jsonInterface struct {
	Text string `json:"text"`
}

type catFact struct {
	Text string `json:"text"`
}

func (core *CoreProcessor) healthCheck(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "OK")
}

func (core *CoreProcessor) apiHandler(resp http.ResponseWriter, req *http.Request) {
	var fact string
	var err error

	if core.offlineMode {
		fact, err = core.catFactOffline()
	} else {
		fact, err = core.catFactOnline()
	}

	if err != nil {
		core.log.Error("Cat API request failed")
	} else {
		core.log.Debug("Fact request successfully processed")
		fmt.Fprintf(resp, "%v", fact)
	}
}

func (core *CoreProcessor) catFactOnline() (string, error) {
	url := "https://cat-fact.herokuapp.com/facts/random"
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		core.log.Errorf("Failed to send HTTP request %+v", err)
		return "", err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var v catFact
	err = decoder.Decode(&v)
	if err != nil {
		core.log.Errorf("Failed to decode HTTP response %+v", err)
		return "", err
	}

	jsonString := &jsonInterface{Text: v.Text}
	b, err := json.Marshal(jsonString)
	if err != nil {
		core.log.Errorf("Failed to marshal text %+v", err)
		return "", err
	}

	return string(b), nil
}

func (core *CoreProcessor) catFactOffline() (string, error) {
	core.log.Debug("Run offline mode")
	var offlineCatFact = []string{
		`Cat's cannot see in total darkness, however their vision is much better than 
		a human's in semidarkness because their retinas are much more sensitive to light.`,
		`Mountain lions are strong jumpers, thanks to muscular hind legs that are longer than their front legs.`,
		`Cats lose almost as much fluid in the saliva while grooming themselves as they do through urination.`,
		`When your cats rubs up against you, she is actually marking you as "hers" with her scent. \
		If your cat pushes his face against your head, it is a sign of acceptance and affection.`,
		`People who own cats have on average 2.1 pets per household, where dog owners have about 1.6.`,
		`After humans, mountain lions have the largest range of any mammal in the Western Hemisphere.`,
	}

	jsonString := &jsonInterface{Text: offlineCatFact[rand.Intn(len(offlineCatFact))]}
	b, err := json.Marshal(jsonString)
	if err != nil {
		core.log.Errorf("Failed to marshal text %+v", err)
		return "", err
	}

	return string(b), err
}

// Start method
func (core *CoreProcessor) Start() {
	core.log.Info("Start factfinder service")

	http.HandleFunc("/health", core.healthCheck)
	http.HandleFunc("/factfinder", core.apiHandler)

	localPort := fmt.Sprintf(":%d", core.port)

	core.log.Fatal(http.ListenAndServe(localPort, nil))

}

// Stop method
func (core *CoreProcessor) Stop() {
	core.log.Info("Stop factfinder service")
}
