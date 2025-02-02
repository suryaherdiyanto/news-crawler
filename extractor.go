package main

import (
	"net/url"

	"golang.org/x/net/html"
)

type LinkExtractor interface {
	GetArticleTags(node *html.Node) []html.Node
	GetLinkText(node *html.Node) string
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

func GetNewsLinks(doc *html.Node, extractor LinkExtractor) []NewsLink {

	var links []NewsLink

	articles := extractor.GetArticleTags(doc)
	for _, articleNode := range articles {
		anchorNodes := GetTags(&articleNode, "a")

		for _, anchorNode := range anchorNodes {
			attr, ok := GetAttribute(&anchorNode, "href")
			text := extractor.GetLinkText(&anchorNode)

			if !ok || attr.Val == "#" {
				continue
			}

			if text == "" {
				continue
			}

			links = append(links, NewsLink{Title: text, Url: attr.Val})
		}
	}

	return links
}
