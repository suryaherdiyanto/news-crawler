package main

import (
	"golang.org/x/net/html"
)

func GetNewsLinks(doc *html.Node) []NewsLink {

	var links []NewsLink

	articles := GetTags(doc, "article")
	for _, articleNode := range articles {
		anchorNodes := GetTags(&articleNode, "a")

		for _, anchorNode := range anchorNodes {
			attr, ok := GetAttribute(&anchorNode, "href")
			text, okText := GetText(&anchorNode)

			if !ok || attr.Val == "#" {
				continue
			}

			if !okText {
				continue
			}

			links = append(links, NewsLink{Title: text, Url: attr.Val})
		}
	}

	return links
}
