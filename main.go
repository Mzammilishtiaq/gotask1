package main

import (
	"fmt"
	"os"
)

// WordCount counts the number of words in a string.
func WordCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word++
		}
	}
	return word
}
func main() {
	fmt.Println("Hello, World!")
	filename, err := os.ReadFile("text_file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(filename)
	totalWords := WordCount(str)
	fmt.Println("Total words in file:", totalWords)
}
