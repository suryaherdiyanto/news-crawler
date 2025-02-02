package main

import (
	"strings"

	"golang.org/x/net/html"
)

func TransverseDecendants(doc *html.Node, fn func(doc *html.Node)) {
	for node := doc.FirstChild; node != nil; node = node.NextSibling {
		fn(node)
		if node.Type != html.ErrorNode {
			TransverseDecendants(node, fn)
		}
	}
}

func GetTags(doc *html.Node, tagName string) []html.Node {
	var nodes []html.Node

	TransverseDecendants(doc, func(node *html.Node) {
		if node.Type == html.ElementNode && node.DataAtom.String() == tagName {
			nodes = append(nodes, *node)
		}
	})

	return nodes
}

func GetAttribute(node *html.Node, name string) (html.Attribute, bool) {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr, true
		}
	}

	return html.Attribute{}, false
}

func GetText(node *html.Node) (string, bool) {
	text := ""

	TransverseDecendants(node, func(node *html.Node) {
		if node.Type == html.TextNode {
			text = node.Data
		}
	})

	text = strings.TrimSpace(text)

	if text == "" {
		return "", false
	}

	return text, true
}

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
