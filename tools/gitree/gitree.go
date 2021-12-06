package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	fmt.Printf("dir=%s\n", dir)

	if !PathExists(dir) {
		fmt.Println("dir not exists")
		return
	}

	Gitree(dir)
}

func Gitree(dir string) {
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && PathExists(filepath.Join(path, ".git")) {
			name, _ := filepath.Rel(dir, path)
			fmt.Printf("\n%s\n", name)
			if urls, err := GitUrls(path); err == nil {
				fmt.Println("  urls:")
				for _, url := range urls {
					if url.tag == "push" {
						fmt.Printf("    %s=%s\n", url.name, url.url)
					}
				}
			} else {
				fmt.Printf("  err=%s\n", err)
			}
			return filepath.SkipDir
		}
		return nil
	})
}

type GitUrl struct {
	name string
	url  string
	tag  string
}

func GitUrls(dir string) ([]GitUrl, error) {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	var urls []GitUrl
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		fields_n := len(fields)
		if fields_n >= 2 {
			url := GitUrl{fields[0], fields[1], ""}
			if fields_n >= 3 {
				url.tag = strings.Trim(fields[2], "()")
			}
			urls = append(urls, url)
		}
	}
	return urls, nil
}

func PathExists(name string) bool {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
