package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filename := "data/data.txt"
	{ // test
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Printf("Not exist: %s\n", filename)

			// create dir if not exist
			dir := filepath.Dir(filename)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			}
		} else {
			log.Printf("Exist: %s\n", filename)
		}
	}
	{ // write
		message := []byte("I love golang!")

		log.Printf("Write file: %s\n", filename)
		log.Printf("  contents: %s\n", message)

		err := ioutil.WriteFile(filename, message, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	{ // append
		message := []byte(" Go coding!")

		log.Printf("Append file: %s\n", filename)
		log.Printf("   contents: %s\n", message)

		err := appendFile(filename, message, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	{ // read
		log.Printf("Read file: %s\n", filename)

		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("  contents: %s", content)
	}
}

func appendFile(filename string, message []byte, perm fs.FileMode) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	if _, err := f.Write(message); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
