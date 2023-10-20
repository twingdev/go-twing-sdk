package common

import (
	"context"
	"encoding/json"
	"time"
)

var throw = ExceptionManager

type Configuration struct {
	maxGramSize     int
	keepTermText    bool
	localPrivateKey []byte
	cache           *InMemoryCache
	netCfg          *NetworkConfiguration
}

type NetworkConfiguration struct {
	ctx              context.Context
	Name             string `json:"net_name"`
	Version          string `json:"net_version"`
	time             time.Time
	timeout          time.Duration
	maxIncoming      int
	maxOutgoing      int
	minThresholdConn int
	validators       map[string]interface{}
	bootstrapNodes   []string
	Hostname         string `json:"net_hostname"`
	Port             string `json:"net_port"`
	keySeed          []byte
	machineID        []byte
	NetAddress       []byte `json:"net_internal_address"`
}

func NewConfiguration(maxGram int, keepTermText bool) *Configuration {

	cache := TwingCache
	return &Configuration{
		maxGramSize:     maxGram,
		keepTermText:    keepTermText,
		localPrivateKey: make([]byte, 0),
		cache:           cache,
	}

}

func initConfiguration() {
	config := NewConfiguration(10, true)
	configBytes, _ := json.Marshal(config)
	err := LocalCache.Set(context.Background(), "config", configBytes)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
}

func GetConfigFromCache() *Configuration {
	var c Configuration
	configBytes, _ := LocalCache.Get(context.Background(), "config")
	_ = json.Unmarshal(configBytes, &c)
	return &c
}

func ReinitializeConfiguration() error {
	err := LocalCache.Delete(context.Background(), "config")
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}
	initConfiguration()
	return nil
}
