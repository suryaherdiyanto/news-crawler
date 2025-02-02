package main

import "golang.org/x/net/html"

func TransverseDecendants(doc *html.Node, fn func(doc *html.Node)) {
	for node := doc.FirstChild; node != nil; node = doc.NextSibling {
		fn(node)
		if node.Type != html.ErrorNode {
			TransverseDecendants(node, fn)
		}
	}
}

func GetTags(doc *html.Node, tagName string) []html.Node {
	var nodes []html.Node

	TransverseDecendants(doc, func(node *html.Node) {
		if node.Type == html.ElementNode && node.Namespace == tagName {
			nodes = append(nodes, *node)
		}
	})

	return nodes
}

func GetNewsLinks(doc *html.Node) []NewsLink {

	var links []NewsLink

	return links
}
