package main

import (
	"fmt"
	"os"
	"time"
)

// WordCount counts the number of words in a string.
func WordCount(s string, ch chan int) {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word++
		}
	}
	ch <- word
}

func LineCount(s string, ch chan int) {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			word++
		}
	}
	ch <- word
}

func PunctcationCount(s string, ch chan int) {
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
	ch <- word
}

func VowelCount(s string, ch chan int) {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' ||
			s[i] == 'o' || s[i] == 'u' || s[i] == 'A' ||
			s[i] == 'E' || s[i] == 'I' || s[i] == 'O' ||
			s[i] == 'U' {
			word++
		}
	}
	ch <- word
}

func ParaCount(s string, ch chan int) {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\t' || s[i] == '\r' || s[i] == '\f' || s[i] == ' ' {
			word++
		}
	}
	ch <- word
}

func SpecialCharacterCount(s string, ch chan int) {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '@' || s[i] == '#' ||
			s[i] == '$' || s[i] == '%' || s[i] == '^' ||
			s[i] == '&' || s[i] == '*' || s[i] == '+' ||
			s[i] == '-' {
			word++
		}
	}
	ch <- word
}

func DigitsCount(s string, ch chan int) {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' || s[i] == '1' || s[i] == '2' ||
			s[i] == '3' || s[i] == '4' || s[i] == '5' ||
			s[i] == '6' || s[i] == '7' || s[i] == '8' ||
			s[i] == '9' {
			word++
		}
	}
	ch <- word
}

func ConsientCount(s string, ch chan int) {
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
	ch <- word
}

func combinefunction(s string, ch chan int) {
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
	ch <- word
}

func main() {
	start := time.Now()
	fmt.Println("Hello, World!")
	filename, err := os.ReadFile("text_file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(filename)
	// channel create in word count
	wordCountChannel := make(chan int)
	go WordCount(str, wordCountChannel)

	// channel create in line count
	LinecountChannel := make(chan int)
	go LineCount(str, LinecountChannel)

	// channel create in punctcation count
	PunctcationCountChannel := make(chan int)
	go PunctcationCount(str, PunctcationCountChannel)

	// channel create in vowel count
	totalvowelChannel := make(chan int)
	go VowelCount(str, totalvowelChannel)

	// channel create in vowel count
	totalparaChannel := make(chan int)
	go ParaCount(str, totalparaChannel)

	// channel create in vowel count
	totalSpecialCharacterChannel := make(chan int)
	go SpecialCharacterCount(str, totalSpecialCharacterChannel)

	// channel create in digit count
	digitcontchannel := make(chan int)
	go DigitsCount(str, digitcontchannel)

	//channel create in combine function
	totalcombinechannel := make(chan int)
	go combinefunction(str, totalcombinechannel)

	// goroutines
	totalWords := <-wordCountChannel
	totallines := <-LinecountChannel
	TotalPunctcationCount := <-PunctcationCountChannel
	totalVowel := <-totalvowelChannel
	totalParagraph := <-totalparaChannel
	totalSpecialCharacterCount := <-totalSpecialCharacterChannel
	totalDigitCount := <-digitcontchannel
	totalCombinefunction := <-totalcombinechannel

	fmt.Println("Total words in file:", totalWords)
	fmt.Println("Total line in file:", totallines)
	fmt.Println("Total PunctcationCount in file:", TotalPunctcationCount)
	fmt.Println("Total VowelCount in file:", totalVowel)
	fmt.Println("Total ParagraphCount in file:", totalParagraph)
	fmt.Println("Total SpecialCharacterCount in file:", totalSpecialCharacterCount)
	fmt.Println("Total DigitCount in file:", totalDigitCount)
	fmt.Println("Combine function:", totalCombinefunction)

	elapse := time.Since(start)
	fmt.Printf("Total time : %s", elapse)
}
