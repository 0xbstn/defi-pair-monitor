package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Config struct {
	Ethereum string `json:"Ethereum"`
	Arbitrum string `json:"Arbitrum"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfigInstance(filename string) *Config {
	once.Do(func() {
		data, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("Unable to read config file: %v", err)
		}

		var config Config
		err = json.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("Unable to parse config file: %v", err)
		}

		instance = &config
	})
	return instance
}
