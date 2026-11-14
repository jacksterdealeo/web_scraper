package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"time"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func composedTime() string {
	n := time.Now()
	return fmt.Sprintf("%vh:%vm:%vs", n.Hour(), n.Minute(), n.Second())
}

func main() {
	cfg := config{
		pages:              map[string]PageData{},
		baseURL:            &url.URL{},
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}),
		wg:                 &sync.WaitGroup{},
	}
	if len(os.Args) <= 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) >= 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %v\n", os.Args[1])
	/*
		d, err := getHTML(os.Args[1])
		if err != nil {
			log.Println(err)
		}
		fmt.Println(d)
	*/

	cfg.crawlPage(os.Args[1], os.Args[1])
	for page, data := range cfg.pages {
		fmt.Println(page)
		_ = data
	}

}
