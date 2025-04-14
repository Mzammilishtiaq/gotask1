package chunkinganalyzer

import (
	"fmt"
	"os"
	"sync"
	"time"
	// "github.com/Mzammilishtiaq/gotask1/chunkingAnalyzer"
)

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

type Result struct{}

func Chunking(str string) Result {
	combinestarttime := time.Now()

	totalLength := len(str)
	chunkpark := 5
	chunkSize := totalLength / chunkpark
	fmt.Println("Total length of the file:", totalLength)
	fmt.Println("chunck size:", chunkSize)
	Chunks := make([]string, chunkpark)

	// Example usage of combinefunction to avoid unused error
	words, lines, spaces, punct, vowels, digits, consonants, specials := combinefunction(str)
	fmt.Printf("Initial Combine Function Results:\nWords: %d, Lines: %d, Spaces: %d, Punctuation: %d, Vowels: %d, Digits: %d, Consonants: %d, Specials: %d\n", words, lines, spaces, punct, vowels, digits, consonants, specials)

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
    // return Result{}

}

func AnalyzeFunction(filepath string) (Result, error) {

	file, err := os.ReadFile(filepath)
	if err != nil {
		return Result{}, err
	}

	return Chunking(string(file)), nil
}
