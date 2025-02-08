package main

import (
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
	GetUpdatedAt() time.Time
	GetCreatedAt() time.Time
}

type CnnNews struct {
	Url string
}

type CnnNewsContent struct {
	html *html.Node
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
