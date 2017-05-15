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

type cache struct {
	mux        sync.Mutex
	visitedMap map[string]bool
}

type response struct {
	url, body string
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache cache, ch chan response) {
	defer close(ch)

	if depth <= 0 {
		return
	}

	cache.mux.Lock()
	if cache.visitedMap[url] {
		cache.mux.Unlock()
		return
	}
	cache.visitedMap[url] = true
	cache.mux.Unlock()


	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch <- response{url, body}
	result := make([]chan response, len(urls))
	for i, u := range urls {
		result[i] = make(chan response)
		go Crawl(u, depth-1, fetcher, cache, result[i])
	}

	for i := range result {
		for resp := range result[i] {
			ch <- resp
		}
	}

	return
}

func main() {
	var ch = make(chan response)
	c := cache{visitedMap: make(map[string]bool)}
	go Crawl("http://golang.org/", 4, fetcher, c, ch)

	for resp := range ch {
		fmt.Printf("found %s %q\n", resp.url, resp.body)
	}
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
