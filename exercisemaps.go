package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordcount := make(map[string]int)
	for _, v := range words {
		wordcount[v]++
	}
	//return map[string]int{"x": 1}
	return wordcount
}

func main() {
	wc.Test(WordCount)
}

