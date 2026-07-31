package main

import (
	"fmt"
	"strings"
)

func autoComplete(line string, builtins []string) {
	var words []string
	for _, word := range builtins {
		if strings.HasPrefix(word, line) {
			words = append(words, word)
		}
	}
	fmt.Println(words)
}

func main() {
	autoComplete("e", []string{"echo,exit"})
}
