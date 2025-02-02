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

func TestGetAttribute(t *testing.T) {
	f, _ := os.Open("./stub/cnn.html")
	defer f.Close()

	parsed, err := html.Parse(f)

	if err != nil {
		t.Error(err)
	}

	tag := GetTags(parsed, "a")[0]
	attr, ok := GetAttribute(&tag, "href")

	if !ok {
		t.Errorf("Expected true, but got %v", ok)
	}

	if attr.Val == "" {
		t.Errorf("Expected to be not empty")
	}
}
