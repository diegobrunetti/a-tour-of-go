package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	
	for _, word := range strings.Fields(s) {
		wordCount[word]++
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
