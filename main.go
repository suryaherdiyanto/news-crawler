package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	page, err := strconv.Atoi(os.Args[2])

	if err != nil {
		panic(err)
	}

	extractor := NewLinkExtractor(url)

	fmt.Println("Fetching: " + url)
	var links []NewsLink
	wg := sync.WaitGroup{}

	for i := 1; i <= page; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond + time.Duration(i*50))
		if i == 1 {
			url = url + "?page=1"
		}

		if i > 1 {
			url = strings.Replace(url, "?page="+strconv.Itoa(i-1), "?page="+strconv.Itoa(i), 1)
		}

		go func(url string) {
			request, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			if request.StatusCode != 200 {
				panic(errors.New("Unexpected status code : " + request.Status))
			}

			fmt.Printf("Crawled <%s> Status code: %d\n", url, request.StatusCode)

			doc, err := html.Parse(request.Body)
			links = append(links, GetNewsLinks(doc, extractor)...)
			wg.Done()
		}(url)
	}

	wg.Wait()

	fmt.Println("done")
	fmt.Printf("Total links found: %d\n", len(links))
}
