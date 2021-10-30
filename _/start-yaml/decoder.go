package main

import (
	"log"
	"os"

	"github.com/ikuokuo/start-golang/_/start-yaml/common"
	"gopkg.in/yaml.v2"
)

func main() {
	filename := "config.yaml"
	log.Printf("open file: %s\n", filename)
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // ignore error

	dec := yaml.NewDecoder(f)

	var cfg common.Config
	if err := dec.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
	log.Printf("decode success! 🍺\n  %v\n", cfg)
}
