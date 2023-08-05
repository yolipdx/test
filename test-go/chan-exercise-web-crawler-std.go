package main

import (
	"fmt"
	// "sync"
)

func Crawl(url string, depth int, fetcher Fetcher, ret chan string) {
    defer close(ret)
    if depth <= 0 {
        return
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        ret <- err.Error()
        return
    }

    ret <- fmt.Sprintf("found: %s %q", url, body)

    result := make([]chan string, len(urls))
    for i, u := range urls {
        result[i] = make(chan string)
        go Crawl(u, depth-1, fetcher, result[i])
    }

    for i := range result {
        for s := range result[i] {
            ret <- s
        }
    }

    return
}

func main() {
    result := make(chan string)
    go Crawl("http://golang.org/", 4, fetcher, result)

    for s := range result {
        fmt.Println(s)
    }
}