package main

import (
	"slices"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type CnnNewsContent struct {
	html *html.Node
}

func (c *CnnNewsContent) GetTitle() string {
	tag := GetTags(c.html, "title")[0]
	title, _ := GetText(&tag)

	return title
}

func (c *CnnNewsContent) GetExcerpt() string {
	return GetMeta(c.html, "description")
}

func (c *CnnNewsContent) GetContent() string {
	divs := GetTags(c.html, "div")
	var paragraphs []html.Node

	for _, div := range divs {
		attr, _ := GetAttribute(&div, "class")
		classList := strings.Split(attr.Val, " ")

		if slices.Index(classList, "data-text") > -1 {
			paragraphs = GetTags(&div, "p")
			break
		}
	}

	var content string

	for _, p := range paragraphs {
		content += "<p>" + strings.TrimSpace(p.Data) + "</p>"
	}

	return content
}

func (c *CnnNewsContent) GetThumbnail() string {
	tags := GetTags(c.html, "div")
	var imageWrapper *html.Node
	
	for _, tag := range tags {
		attr, _ := GetAttribute(&tag, "class")
		classList := strings.Split(attr.Val, " ")

		if slices.Index(classList, "detail-image") > -1 {
			imageWrapper = &tag
		}
	}

	if imageWrapper.DataAtom.String() == "" {
		return ""
	}

	img := GetTags(imageWrapper, "img")[0]
	attr, _ := GetAttribute(&img, "src")

	return attr.Val
}

func (c *CnnNewsContent) GetPublishedAt() time.Time {
	publishDate := GetMeta(c.html, "dtk:createddate")
	t, err := time.Parse("2006/01/02 15:04:05 MST", publishDate + " WIB")

	if err != nil {
		panic(err)
	}

	return t
}

func (c *CnnNewsContent) GetCreatedAt() time.Time {
	createdDate := GetMeta(c.html, "dtk:createddate")
	t, err := time.Parse("2006/01/02 15:04:05 MST", createdDate + " WIB")

	if err != nil {
		panic(err)
	}

	return t
}
