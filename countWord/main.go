package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "This is jahidul islam, how are you?"

	words := strings.Split(text, " ")

	fmt.Println("Number of words in text: ", words)
	fmt.Println("Number of words in text: ", len(words))
}
