package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlContext struct {
	maxdepth  int
	taskCount map[string]int
	taskMemo  map[string]bool
	taskMux   sync.Mutex
}

func (c *CrawlContext) TaskInc(url string) int {
	c.taskMux.Lock()
	defer c.taskMux.Unlock()
	n, ok := c.taskCount[url]
	if ok {
		n += 1
	} else {
		n = 1
	}
	c.taskCount[url] = n
	return n
}

func (c *CrawlContext) TaskDec(url string) int {
	c.taskMux.Lock()
	defer c.taskMux.Unlock()
	n, ok := c.taskCount[url]
	if ok {
		n -= 1
		if n == 0 {
			delete(c.taskCount, url)
		} else {
			c.taskCount[url] = n
		}
		return n
	} else {
		return 0
	}
}

func (c *CrawlContext) TaskAllDone() bool {
	c.taskMux.Lock()
	defer c.taskMux.Unlock()
	return len(c.taskCount) == 0
}

func (c *CrawlContext) TaskMemo(url string, result bool) {
	c.taskMux.Lock()
	c.taskMemo[url] = result
	c.taskMux.Unlock()
}

func (c *CrawlContext) TaskExist(url string) bool {
	c.taskMux.Lock()
	defer c.taskMux.Unlock()
	_, ok := c.taskMemo[url]
	return ok
}

func crawl(url string, depth int, fetcher Fetcher, ctx *CrawlContext, quit chan int) {
	if depth == 0 {
		ctx.TaskInc(url)
	}
	defer func() {
		ctx.TaskDec(url)
		if ctx.TaskAllDone() {
			quit <- 0
		}
	}()

	if depth >= ctx.maxdepth || ctx.TaskExist(url) {
		return
	}

	ctx.TaskMemo(url, false)
	body, urls, err := fetcher.Fetch(url)
	ctx.TaskMemo(url, true)

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, u := range urls {
		ctx.TaskInc(u)
		go crawl(u, depth+1, fetcher, ctx, quit)
	}

	// Process
	fmt.Printf("found: %s %q\n", url, body)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	quit := make(chan int)
	ctx := &CrawlContext{
		maxdepth:  depth,
		taskCount: make(map[string]int),
		taskMemo:  make(map[string]bool)}
	go crawl(url, 0, fetcher, ctx, quit)
	<-quit
	fmt.Println("quit")
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
