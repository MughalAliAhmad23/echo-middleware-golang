package filereader

import (
	"calculator/models"
	"fmt"
	"sync"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func CheckByteInSlice(b byte, slice []byte) bool {
	for i := 0; i < len(slice); i++ {
		if b == slice[i] {
			return true
		}
	}
	return false
}

func VowelsCheck(b byte) bool {
	vowels := "aeiouAEIOU"
	return CheckByteInSlice(b, []byte(vowels))
}

func PunctuationCheck(b byte) bool {
	punctuation := ",?!:;—-()[]{}'\".../\\<>_&*^~`|"
	return CheckByteInSlice(b, []byte(punctuation))
}

func FileProcessor(chunk []byte, chanResult chan<- models.Filestats, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		lineCount        = 0
		wordCount        = 0
		vowelsCount      = 0
		punctuationCount = 0
		atWord           = false
	)

	for _, val := range chunk {
		switch {
		case val == '\n':
			lineCount++
		case val == ' ' || val == '\t':
			if atWord {
				wordCount++
				atWord = false
			}
		default:
			atWord = true
			if VowelsCheck(val) {
				vowelsCount++
			}
			if PunctuationCheck(val) {
				punctuationCount++
			}
		}
	}
	if atWord {
		wordCount++
	}
	chanResult <- models.Filestats{
		Totallines:       lineCount,
		Totalwords:       wordCount,
		Totalvowels:      vowelsCount,
		Totalpunctuation: punctuationCount,
	}
}
