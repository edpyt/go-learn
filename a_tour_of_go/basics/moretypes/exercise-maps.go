package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	myWordCounter := make(map[string]int)
	sFields := strings.Fields(s)
	
	for i := 0; i < len(sFields); i++ {
		myWordCounter[sFields[i]] += 1
	}

	return myWordCounter
}

func main() {
	wc.Test(WordCount)
}
