package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestGetTags(t *testing.T) {
	f, _ := os.Open("./stub/cnn.html")
	defer f.Close()

	parsed, err := html.Parse(f)

	if err != nil {
		t.Error(err)
	}

	tags := GetTags(parsed, "article")

	if len(tags) == 0 {
		t.Errorf("Expected tags to be more than 0")
	}
}
