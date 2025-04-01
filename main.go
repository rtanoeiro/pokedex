package main

import (
	"fmt"
	"strings"
)

func main() {
	split:=cleanInput("  hello  world  ")
	fmt.Println("Size of split word: ", len(split))
}

func cleanInput(text string) []string {
	
	lowerWord := strings.ToLower(text)
	splitText := strings.Fields(lowerWord)

	return splitText
}