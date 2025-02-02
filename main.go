package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url, err := url.Parse(os.Args[1])

	if err != nil {
		panic(err)
	}

	fmt.Println("Fetching: " + url.String())
	request, err := http.Get(url.String())
	fmt.Println("Fetched, with status code: " + request.Status)

	if request.StatusCode != 200 {
		panic(errors.New("Unexpected status code: " + request.Status))
	}

	doc, err := html.Parse(request.Body)
	links := GetNewsLinks(doc)

	fmt.Println(links)
	fmt.Println(len(links))
}
