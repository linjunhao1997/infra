package config

import (
	"os"

	"github.com/hashicorp/consul/api"
)

func GetConfigFromConsul() []byte {
	endpoints := os.Getenv("CONSUL_ENDPOINTS")
	if len(endpoints) == 0 {
		panic("env CONSUL_ENDPOINTS empty")
	}
	client, err := api.NewClient(&api.Config{Scheme: "http", Address: endpoints})
	if err != nil {
		panic(err)
	}
	kv := client.KV()
	configKey := os.Getenv("CONFIG_KEY")
	if len(configKey) == 0 {
		panic("env CONFIG_KEY empty")
	}
	data, _, err := kv.Get(configKey, nil)
	if err != nil {
		panic(err)
	}
	return data.Value
}
