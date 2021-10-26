package main

import (
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("data/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("File contents: %s", content)
}
