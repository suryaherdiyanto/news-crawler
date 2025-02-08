package main

import (
	"fmt"
	"testing"
)

func TestCnnExtractor(t *testing.T) {
	url := "https://www.cnnindonesia.com/nasional/20250208105015-20-1195986/16-rt-dan-4-ruas-jalan-di-jakarta-terendam-banjir-sabtu-pagi"
	extractor := NewContentExtractor(url)	

	fmt.Println(extractor.GetExcerpt())
	fmt.Println(extractor.GetContent())
}
