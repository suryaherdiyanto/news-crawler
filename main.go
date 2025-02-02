package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]

	extractor := NewLinkExtractor(url)

	fmt.Println("Fetching: " + url)
	request, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Fetched, with status code: " + request.Status)

	if request.StatusCode != 200 {
		panic(errors.New("Unexpected status code: " + request.Status))
	}

	doc, err := html.Parse(request.Body)
	links := GetNewsLinks(doc, extractor)

	fmt.Println(links)
	fmt.Println(len(links))
}
