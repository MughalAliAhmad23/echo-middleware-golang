package filereader

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func SpaceCounter(s string, wg *sync.WaitGroup, c chan<- string) {
	defer wg.Done()
	spaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaces++
		}
	}
	resp := fmt.Sprintf("%d totalspaces:", spaces)
	fmt.Println(resp)
	c <- resp
}

func VowelsCounter(s string, wg *sync.WaitGroup, c chan<- string) {
	defer wg.Done()
	vowels := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowels++
		}

	}
	resp := fmt.Sprintf("%d total vowels:", vowels)
	fmt.Println(resp)
	c <- resp
}

func LineCounter(s string, wg *sync.WaitGroup, c chan<- string) {
	defer wg.Done()
	lines := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			lines++
		}
	}
	resp := fmt.Sprintf("%d total lines :", lines)
	fmt.Println(resp)
	c <- resp
}

func Wordfrequeny(s string, wg *sync.WaitGroup, c chan<- string) {
	defer wg.Done()

	wordcount := make(map[string]int)

	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		if value, key := wordcount[split[i]]; key {
			wordcount[split[i]] = value + 1
		} else {
			wordcount[split[i]] = 1
		}
	}

	resp := fmt.Sprintf("%v word frequencies:", wordcount)
	fmt.Println(resp)
	c <- resp
}

func Wordcounter(s string, wg *sync.WaitGroup, c chan<- string) {

	defer wg.Done()

	spaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaces++
		}
	}
	totalwords := spaces + 1

	resp := fmt.Sprintf("%d total words:", totalwords)
	fmt.Println(resp)
	c <- resp
}
