package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	Lower := strings.ToLower(text)
	split_words := strings.Fields(Lower)
	return split_words
}

func main() {
	fmt.Println("Hello, World!")
}
