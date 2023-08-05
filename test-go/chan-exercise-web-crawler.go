package main

import (
	"fmt"
	"sync"
)

type mark struct {
	contents map[string]bool
	mutex    sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, mark *mark) {
	done := make(chan int)
	go DoCrawl(url, depth, fetcher, mark, done)
	<-done
}

func DoCrawl(url string, depth int, fetcher Fetcher, mark *mark, done chan int) {
	defer func() {
		done <- 1
	}()

	mark.mutex.Lock()
	if _, ok := mark.contents[url]; ok {
		// TODO: Don't fetch the same URL twice.
		fmt.Println("ignore: ", url)
		defer mark.mutex.Unlock()
		return
	}
	mark.mutex.Unlock()

	mark.mutex.Lock()
	mark.contents[url] = true
	mark.mutex.Unlock()
	// TODO: Fetch URLs in parallel.

	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	
	doneChan := make(chan int)
	for _, u := range urls {
		go DoCrawl(u, depth-1, fetcher, mark, doneChan)
	}

	for i:=0; i <len(urls); i++ {
		<-doneChan
	}

}

func main() {
	mark := mark{
		contents: make(map[string]bool),
	}
	Crawl("https://golang.org/", 4, fetcher, &mark)
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
