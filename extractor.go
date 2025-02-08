package main

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

type LinkExtractor interface {
	GetArticleTags(node *html.Node) []html.Node
	GetLinkText(node *html.Node) string
}

type ContentExtractor interface {
	GetTitle() string
	GetExcerpt() string
	GetContent() string
	GetThumbnail() string
	GetPublishedAt() time.Time
	GetCreatedAt() time.Time
}

type CnnNews struct {
	Url string
}

func (c *CnnNews) GetArticleTags(node *html.Node) []html.Node {
	return GetTags(node, "article")
}

func (c *CnnNews) GetLinkText(node *html.Node) string {
	text, _ := GetTextFromChilds(node, "h2")
	return text
}


func NewLinkExtractor(u string) LinkExtractor {
	parsedUrl, err := url.Parse(u)

	if err != nil {
		panic(err)
	}

	switch parsedUrl.Host {
	case "www.cnnindonesia.com":
		return &CnnNews{Url: parsedUrl.String()}
	default:
		return &CnnNews{}
	}
}

func NewContentExtractor(u string) ContentExtractor {
	parsedUrl, err := url.Parse(u)

	if err != nil {
		panic(err)
	}

	res, err := http.Get(parsedUrl.String())

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic(errors.New("Invalid status code: " + res.Status))
	}

	doc, err := html.Parse(res.Body)

	if err != nil {
		panic(err)
	}

	switch parsedUrl.Host {
	case "www.cnnindonesia.com":
		return &CnnNewsContent{html: doc}
	default:
		return &CnnNewsContent{}
	}
}
