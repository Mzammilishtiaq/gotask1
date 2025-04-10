package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// WordCount counts the number of words in a string.
// func WordCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == ' ' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

// func LineCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '\n' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

// func PunctcationCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '.' || s[i] == '?' || s[i] == '!' || s[i] == ',' ||
// 			s[i] == ';' ||
// 			s[i] == ':' || s[i] == '"' ||
// 			s[i] == '\'' || s[i] == '(' || s[i] == ')' ||
// 			s[i] == '[' || s[i] == ']' || s[i] == '{' ||
// 			s[i] == '}' || s[i] == '<' || s[i] == '>' ||
// 			s[i] == '/' || s[i] == '\\' || s[i] == '^' ||
// 			s[i] == '_' || s[i] == '=' || s[i] == '|' ||
// 			s[i] == '`' || s[i] == '~' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

// func VowelCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' ||
// 			s[i] == 'o' || s[i] == 'u' || s[i] == 'A' ||
// 			s[i] == 'E' || s[i] == 'I' || s[i] == 'O' ||
// 			s[i] == 'U' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

// func ParaCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '\n' || s[i] == '\t' || s[i] == '\r' || s[i] == '\f' || s[i] == ' ' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

// func SpecialCharacterCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '@' || s[i] == '#' ||
// 			s[i] == '$' || s[i] == '%' || s[i] == '^' ||
// 			s[i] == '&' || s[i] == '*' || s[i] == '+' ||
// 			s[i] == '-' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

// func DigitsCount(s string, ch chan int) {
// 	word := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '0' || s[i] == '1' || s[i] == '2' ||
// 			s[i] == '3' || s[i] == '4' || s[i] == '5' ||
// 			s[i] == '6' || s[i] == '7' || s[i] == '8' ||
// 			s[i] == '9' {
// 			word++
// 		}
// 	}
// 	ch <- word
// }

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

func combinefunction(s string) (int, int, int, int, int, int, int, int) {
	words := 0
	line := 0
	spaces := 0
	vowels := 0
	punctcation := 0
	digits := 0
	ConsientCount := 0
	SpecialCharacter := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			words++
		}
		if s[i] == '\n' {
			line++
		}
		if s[i] == '	' {
			spaces++
		}
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		switch s[i] {
		case '.', ',', '\'',
			'(', ')', '!', '?',
			';', ':', '-', '"':
			punctcation++
		}
		switch s[i] {
		case '0', '1', '2', '3',
			'4', '5', '6', '7', '8',
			'9':
			digits++
		}
		if s[i] != 'a' && s[i] != 'e' &&
			s[i] != 'i' && s[i] != 'o' &&
			s[i] != 'u' {
			ConsientCount++
		}
		if !(s[i] >= 'a' && s[i] <= 'z') &&
			!(s[i] >= 'A' && s[i] <= 'Z') &&
			!(s[i] >= '0' && s[i] <= '9') &&
			s[i] != ' ' {
			SpecialCharacter++
		}
	}
	return words, line, spaces, punctcation, vowels, digits, ConsientCount, SpecialCharacter
}

func main() {
	combinestarttime := time.Now()
	filename, err := os.ReadFile("text_file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	str := string(filename)
	totalLength := len(str)
	chunkpark := 5
	chunkSize := totalLength / chunkpark
	fmt.Println("Total length of the file:", totalLength)
	fmt.Println("chunck size:", chunkSize)
	Chunks := make([]string, chunkpark)

	for i := 0; i < 3; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == 3-1 {
			end = totalLength
		}

		Chunks[i] = str[start:end]
	}
	ChunkChannel := make(chan string, chunkpark)
	var wg sync.WaitGroup
	for i, chunk := range Chunks {
		wg.Add(chunkpark)
		go func(i int, chunkingData string) {
			defer wg.Done()
			words, lines, spaces,
				punct, vowels, digits,
				consonants, specials := combinefunction(chunkingData)

			result := fmt.Sprintf("\n===========%d == chuncking==========\n", i+1)
			result += fmt.Sprintf("Word Count: %d\n", words)
			result += fmt.Sprintf("Word Count: %d\n", lines)
			result += fmt.Sprintf("Word Count: %d\n", spaces)
			result += fmt.Sprintf("Word Count: %d\n", punct)
			result += fmt.Sprintf("Word Count: %d\n", vowels)
			result += fmt.Sprintf("Word Count: %d\n", digits)
			result += fmt.Sprintf("Word Count: %d\n", consonants)
			result += fmt.Sprintf("Word Count: %d\n", specials)
			ChunkChannel <- result
		}(i, chunk)
	}
	go func() {
		wg.Wait()
		close(ChunkChannel)
	}()
	for result := range ChunkChannel {
		fmt.Println(result)
	}
	combinrendtime := time.Since(combinestarttime)
	fmt.Printf("Total Combine function time : %s", combinrendtime)

	// channel create in word count
	// wordCountChannel := make(chan int)
	// go WordCount(str, wordCountChannel)

	// channel create in line count
	// LinecountChannel := make(chan int)
	// go LineCount(str, LinecountChannel)

	// channel create in punctcation count
	// PunctcationCountChannel := make(chan int)
	// go PunctcationCount(str, PunctcationCountChannel)

	// channel create in vowel count
	// totalvowelChannel := make(chan int)
	// go VowelCount(str, totalvowelChannel)

	// channel create in vowel count
	// totalparaChannel := make(chan int)
	// go ParaCount(str, totalparaChannel)

	// channel create in vowel count
	// totalSpecialCharacterChannel := make(chan int)
	// go SpecialCharacterCount(str, totalSpecialCharacterChannel)

	// channel create in digit count
	// digitcontchannel := make(chan int)
	// go DigitsCount(str, digitcontchannel)

	//channel create in combine function
	// totalcombinechannel := make(chan int)
	// go combinefunction(str, totalcombinechannel)

	// goroutines
	// totalWords := <-wordCountChannel
	// totallines := <-LinecountChannel
	// TotalPunctcationCount := <-PunctcationCountChannel
	// totalVowel := <-totalvowelChannel
	// totalParagraph := <-totalparaChannel
	// totalSpecialCharacterCount := <-totalSpecialCharacterChannel
	// totalDigitCount := <-digitcontchannel

	// fmt.Println("<====================== start all seperate function run============================>")
	// fmt.Println("Total words in file:", totalWords)
	// fmt.Println("Total line in file:", totallines)
	// fmt.Println("Total PunctcationCount in file:", TotalPunctcationCount)
	// fmt.Println("Total VowelCount in file:", totalVowel)
	// fmt.Println("Total ParagraphCount in file:", totalParagraph)
	// fmt.Println("Total SpecialCharacterCount in file:", totalSpecialCharacterCount)
	// fmt.Println("Total DigitCount in file:", totalDigitCount)
	// speratefunctiontime := time.Since(start)
	// fmt.Printf("Total Sperate function time :%s", speratefunctiontime)
	// fmt.Println("\n<====================== end all seperate function run============================>")
	// totalCombinefunction:= <-totalcombinechannel

	// combinestarttime := time.Now()
	// words, lines, spaces, punctuations, vowels, digits, consonants, specialCharacters := combinefunction(str)
	// fmt.Println("\n<====================== start combine function run============================>")
	// fmt.Println("Combine Total words in file:", words)
	// fmt.Println("Combine Total line in file:", lines)
	// fmt.Println("Combine Total space in file:", spaces)
	// fmt.Println("Combine Total vowels in file:", vowels)
	// fmt.Println("Combine Total punctuations in file:", punctuations)
	// fmt.Println("Combine Total digits in file:", digits)
	// fmt.Println("Combine Total consonants in file:", consonants)
	// fmt.Println("Combine Total specialsCharacters in file:", specialCharacters)
	// combinrendtime := time.Since(combinestarttime)
	// fmt.Printf("Total Combine function time : %s", combinrendtime)
	// fmt.Println("\n<======================end combine function run============================>")

}
