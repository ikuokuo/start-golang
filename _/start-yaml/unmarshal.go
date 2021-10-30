package main

import (
	"log"

	"github.com/ikuokuo/start-golang/_/start-yaml/common"
	"gopkg.in/yaml.v3"
)

var data = `
version: 0.0.0
createdat: 2021-10-30T10:00:00+08:00
labels:
  - go
  - coding
server:
  addr: 0.0.0.0
  port: 8000
`

func main() {
	cfg := common.Config{}

	err := yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(cfg)
}
