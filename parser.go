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

func GetMetaTag(doc *html.Node, tagName string, metaName string) (html.Attribute, bool) {
	metas := GetTags(doc, "meta")
	
	for _, meta := range metas {
		attr, ok := GetAttribute(&meta, tagName)

		if ok && attr.Val == metaName {
			return GetAttribute(&meta, "content")
		}
	}

	return html.Attribute{}, false
}

func GetAttribute(node *html.Node, name string) (html.Attribute, bool) {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr, true
		}
	}

	return html.Attribute{}, false
}

func GetMeta(node *html.Node, val string) string {
	meta, ok := GetMetaTag(node, "name", val)

	if !ok {
		return ""
	}

	return meta.Val
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

func GetTextFromChilds(node *html.Node, tag string) (string, bool) {
	text := ""

	TransverseDecendants(node, func(n *html.Node) {
		if n.Data == tag {
			c := n.FirstChild
			if c.Type == html.TextNode {
				nodeText := strings.TrimSpace(c.Data)
				if nodeText != "" {
					text = nodeText
				}
			}
		}
	})

	if text == "" {
		return "", false
	}

	return text, true
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
