package main

import (
	"fmt"
	"strings"
)

func main() {
	var myLine string = "Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure."
	words := strings.Fields(myLine)
	wordCountMap := make(map[string]int)
	for _, word := range words {
		count, wordExists := wordCountMap[word]
		if wordExists {
			wordCountMap[word] = count + 1
		} else {
			wordCountMap[word] = 1
		}
	}
	for key, value := range wordCountMap {
		fmt.Println(key, value)
	}
}
