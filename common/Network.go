package common

import (
	"context"
	"encoding/json"
	"time"
)

type NetworkCfg struct {
	config *NetworkConfiguration
}

func (n *NetworkCfg) Configure(name, version, hostname, port, internalAddress string) {

	n.config.time = time.Now()
	n.config.timeout = time.Second * 60
	n.config.maxIncoming = 100
	n.config.maxOutgoing = 50
	n.config.minThresholdConn = 2
	n.config.validators = make(map[string]interface{})
	n.config.bootstrapNodes = []string{}
	n.config.Hostname = hostname
	n.config.Name = name
	n.config.Version = version
	n.config.Port = port
	n.config.NetAddress = []byte(internalAddress)
	err := LocalCache.Delete(n.config.ctx, "netConfig")
	if err != nil {
		throw.ThrowRecoverableError(err.Error())
	}
	nConfigBytes, _ := json.Marshal(n.config)
	err = LocalCache.Set(n.config.ctx, "netConfig", nConfigBytes)
	if err != nil {
		throw.ThrowFatalError(err.Error())
	}

}

func GetNetworkConfiguration() *NetworkConfiguration {
	nConfigBytes, _ := LocalCache.Get(context.Background(), "nConfig")
	var nc NetworkConfiguration
	json.Unmarshal(nConfigBytes, &nc)
	return &nc
}

func NewNetworkConfiguration(name, version, hostname, port, internalAddress string) {
	n := new(NetworkCfg)
	n.Configure(name, version, hostname, port, internalAddress)
}
