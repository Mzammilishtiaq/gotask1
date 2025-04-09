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

func LineCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			word++
		}
	}
	return word
}

func PunctcationCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '?' || s[i] == '!' || s[i] == ',' ||
			s[i] == ';' ||
			s[i] == ':' || s[i] == '"' ||
			s[i] == '\'' || s[i] == '(' || s[i] == ')' ||
			s[i] == '[' || s[i] == ']' || s[i] == '{' ||
			s[i] == '}' || s[i] == '<' || s[i] == '>' ||
			s[i] == '/' || s[i] == '\\' || s[i] == '^' ||
			s[i] == '_' || s[i] == '=' || s[i] == '|' ||
			s[i] == '`' || s[i] == '~' {
			word++
		}
	}
	return word
}

func VowelCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' ||
			s[i] == 'o' || s[i] == 'u' || s[i] == 'A' ||
			s[i] == 'E' || s[i] == 'I' || s[i] == 'O' ||
			s[i] == 'U' {
			word++
		}
	}
	return word
}

func ParaCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\t' || s[i] == '\r' || s[i] == '\f' || s[i] == ' ' {
			word++
		}
	}
	return word
}

func SpecialCharacterCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '@' || s[i] == '#' ||
			s[i] == '$' || s[i] == '%' || s[i] == '^' ||
			s[i] == '&' || s[i] == '*' || s[i] == '+' ||
			s[i] == '-' {
			word++
		}
	}
	return word
}

func DigitsCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' || s[i] == '1' || s[i] == '2' ||
			s[i] == '3' || s[i] == '4' || s[i] == '5' ||
			s[i] == '6' || s[i] == '7' || s[i] == '8' ||
			s[i] == '9' {
			word++
		}
	}
	return word
}

func ConsientCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' || s[i] == 'c' ||
			s[i] == 'd' || s[i] == 'f' ||
			s[i] == 'g' || s[i] == 'h' ||
			s[i] == 'j' || s[i] == 'k' ||
			s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' ||
			s[i] == 'q' || s[i] == 'r' ||
			s[i] == 's' || s[i] == 'v' ||
			s[i] == 'w' || s[i] == 'x' ||
			s[i] == 'y' || s[i] == 'z' {
			word++
		}
	}
	return word
}

func combinefunction(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '.' ||
			s[i] == '?' || s[i] == '!' || s[i] == ',' ||
			s[i] == ';' || s[i] == ':' || s[i] == '"' ||
			s[i] == '\'' || s[i] == '(' || s[i] == ')' ||
			s[i] == '[' || s[i] == ']' || s[i] == '{' ||
			s[i] == '}' || s[i] == '<' || s[i] == '>' ||
			s[i] == '/' || s[i] == '\\' || s[i] == '^' ||
			s[i] == '_' || s[i] == '=' || s[i] == '|' ||
			s[i] == '`' || s[i] == '~' || s[i] == 'a' ||
			s[i] == 'e' || s[i] == 'i' ||
			s[i] == 'o' || s[i] == 'u' || s[i] == 'A' ||
			s[i] == 'E' || s[i] == 'I' || s[i] == 'O' ||
			s[i] == 'U' || s[i] == '\n' || s[i] == '\t' ||
			s[i] == '\r' || s[i] == '\f' || s[i] == '@' ||
			s[i] == '#' ||
			s[i] == '$' || s[i] == '%' || s[i] == '^' ||
			s[i] == '&' || s[i] == '*' || s[i] == '+' ||
			s[i] == '-' || s[i] == '0' || s[i] == '1' || s[i] == '2' ||
			s[i] == '3' || s[i] == '4' || s[i] == '5' ||
			s[i] == '6' || s[i] == '7' || s[i] == '8' ||
			s[i] == '9' || s[i] == 'b' || s[i] == 'c' ||
			s[i] == 'd' || s[i] == 'f' ||
			s[i] == 'g' || s[i] == 'h' ||
			s[i] == 'j' || s[i] == 'k' ||
			s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' ||
			s[i] == 'q' || s[i] == 'r' ||
			s[i] == 's' || s[i] == 'v' ||
			s[i] == 'w' || s[i] == 'x' ||
			s[i] == 'y' || s[i] == 'z' {
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
	totallines := LineCount(str)
	TotalPunctcationCount := PunctcationCount(str)
	totalVowel := VowelCount(str)
	totalParagraph := ParaCount(str)
	totalSpecialCharacterCount := SpecialCharacterCount(str)
	totalDigitCount := DigitsCount(str)
	totalCombinefunction := combinefunction(str)
	fmt.Println("Total words in file:", totalWords)
	fmt.Println("Total line in file:", totallines)
	fmt.Println("Total PunctcationCount in file:", TotalPunctcationCount)
	fmt.Println("Total VowelCount in file:", totalVowel)
	fmt.Println("Total ParagraphCount in file:", totalParagraph)
	fmt.Println("Total SpecialCharacterCount in file:", totalSpecialCharacterCount)
	fmt.Println("Total DigitCount in file:", totalDigitCount)
	fmt.Println("Combine function:", totalCombinefunction)
}
