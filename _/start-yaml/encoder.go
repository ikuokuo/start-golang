package main

import (
	"log"
	"os"
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

	filename := "config.yaml"
	log.Printf("open file: %s\n", filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // ignore error

	enc := yaml.NewEncoder(f)
	if err := enc.Encode(cfg); err != nil {
		log.Fatal(err)
	}
	log.Println("encode success! 🍺")
}
