package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)
	fmt.Println(words)
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
	//
	return map[string]int{words[0]: len(words)}
}

func main() {

	WordCount("test")
	wc.Test(WordCount)
}
