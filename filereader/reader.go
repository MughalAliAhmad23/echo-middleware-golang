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

// func SpaceCounter(s string, wg *sync.WaitGroup, c chan<- string, a []string) {
// 	defer wg.Done()

// 	var iwg sync.WaitGroup
// 	totalspaces := 0
// 	space := make([]int, len(a))

// 	for j := 0; j < len(a); j++ {
// 		iwg.Add(1)

// 		go func(s string, iwg *sync.WaitGroup, ind int) {
// 			defer iwg.Done()

// 			for i := 0; i < len(s); i++ {
// 				if s[i] == ' ' {
// 					space[ind]++
// 				}
// 			}
// 		}(a[j], &iwg, j)
// 	}
// 	iwg.Wait()

// 	for _, count := range space {
// 		totalspaces += count
// 	}

// 	resp := fmt.Sprintf("%d totalspaces:", totalspaces)
// 	fmt.Println(resp)
// 	c <- resp
// }

// func VowelsCounter(s string, wg *sync.WaitGroup, c chan<- string, a []string) {
// 	defer wg.Done()

// 	totalvowels := 0
// 	var iwg sync.WaitGroup
// 	vowels := make([]int, len(a))

// 	for i := 0; i < len(a); i++ {
// 		iwg.Add(1)

// 		go func(s string, iwg *sync.WaitGroup, ind int) {
// 			defer iwg.Done()

// 			for i := 0; i < len(s); i++ {
// 				if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
// 					vowels[ind]++
// 				}
// 			}

// 		}(a[i], &iwg, i)
// 	}
// 	iwg.Wait()

// 	for _, count := range vowels {
// 		totalvowels += count
// 	}

// 	resp := fmt.Sprintf("%d total vowels:", totalvowels)
// 	fmt.Println(resp)
// 	c <- resp
// }

// func Wordfrequeny(s string, wg *sync.WaitGroup, c chan<- string, a []string) {
// 	defer wg.Done()

// 	var iwg sync.WaitGroup
// 	wordscount := make([]map[string]int, len(a))
// 	temp := make([]string, len(a))

// 	for i := range wordscount {
// 		wordscount[i] = make(map[string]int)
// 	}

// 	for i := 0; i < len(a); i++ {
// 		iwg.Add(1)

// 		go func(s string, iwg *sync.WaitGroup, ind int) {
// 			defer iwg.Done()

// 			for j := 0; j < len(s); j++ {

// 				if s[j] == ' ' || j == len(s)-1 {

// 					if val, key := wordscount[ind][temp[ind]]; key {
// 						wordscount[ind][temp[ind]] = val + 1
// 					} else {
// 						wordscount[ind][temp[ind]] = 1
// 					}
// 					temp[ind] = ""
// 				} else {
// 					temp[ind] += string(s[j])
// 				}
// 			}

// 		}(a[i], &iwg, i)

// 	}
// 	iwg.Wait()

// 	mergemap := make(map[string]int)

// 	for i := 0; i < len(wordscount); i++ {

// 		for key := range wordscount[i] {

// 			if _, ok := mergemap[key]; ok {

// 				mergemap[key] = mergemap[key] + wordscount[i][key]
// 			} else {
// 				mergemap[key] = wordscount[i][key]
// 			}

// 		}
// 	}

// 	resp := fmt.Sprintf("%v word frequencies:", mergemap)
// 	fmt.Println(resp)
// 	c <- resp
// }

// func Wordcounter(s string, wg *sync.WaitGroup, c chan<- string, a []string) {

// 	defer wg.Done()

// 	totalwords := 0
// 	var iwg sync.WaitGroup
// 	words := make([]int, len(a))

// 	for i := 0; i < len(a); i++ {
// 		iwg.Add(1)

// 		go func(s string, iwg *sync.WaitGroup, ind int) {
// 			defer iwg.Done()

// 			for j := 0; j < len(s); j++ {
// 				if s[j] == ' ' || j == len(s)-1 {
// 					words[ind]++
// 				}
// 			}

// 		}(a[i], &iwg, i)
// 	}
// 	iwg.Wait()

// 	for _, count := range words {
// 		totalwords += count
// 	}

// 	resp := fmt.Sprintf("%d total words:", totalwords)
// 	fmt.Println(resp)
// 	c <- resp

// }

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
