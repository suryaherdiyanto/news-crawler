package main

import "time"

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
