package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts
}

func main() {
	wc.Test(WordCount)
}

//Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
//You might find strings.Fields helpful.
