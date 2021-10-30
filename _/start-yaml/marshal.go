package main

import (
	"log"
	"time"

	"github.com/ikuokuo/start-golang/_/start-yaml/common"
	"gopkg.in/yaml.v3"
)

func main() {
	cfg := common.Config{
		V:         "1.0.0",
		CreatedAt: time.Now(),
		Labels:    []string{"go", "coding"},
		Server: common.ServerConfig{
			Addr: "0.0.0.0",
			Port: 8000,
		},
	}

	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Print(string(d))
}
