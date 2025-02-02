package main

import (
	"strings"
	"time"

	"golang.org/x/net/html"
)

type NewsLink struct {
	Title, Url string
}

type NewsArticle struct {
	Title, Excerpt, Content string
	ThumbnailUrl            interface{}
	PublishedAt             time.Time
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

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

	TransverseDecendants(node, func(n *html.Node) {
		if n.Type == html.TextNode {
			nodeText := strings.TrimSpace(n.Data)
			if nodeText != "" {
				text = nodeText
			}
		}
	})

	if text == "" {
		return "", false
	}

	return text, true
}
