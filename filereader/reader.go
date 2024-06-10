package filereader

import (
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

func SpaceCounter(s string, wg *sync.WaitGroup, c chan<- string, a []string) {
	defer wg.Done()

	var iwg sync.WaitGroup

	totalspaces := 0

	space := make([]int, len(a))
	for j := 0; j < len(a); j++ {

		iwg.Add(1)

		go func(s string, iwg *sync.WaitGroup, ind int) {
			//	fmt.Println("index value", space[ind])
			defer iwg.Done()
			//	space[j] := 0

			for i := 0; i < len(s); i++ {
				if s[i] == ' ' {

					space[ind]++

				}
			}
			//fmt.Println("index after addition:", space[ind])
			//totalspaces = +space[ind]
		}(a[j], &iwg, j)
	}
	iwg.Wait()
	//totalspaces := 0
	for _, count := range space {
		totalspaces += count
	}
	//fmt.Println("totalspaces:", totalspaces)

	resp := fmt.Sprintf("%d totalspaces:", totalspaces)
	fmt.Println(resp)
	c <- resp
}

func VowelsCounter(s string, wg *sync.WaitGroup, c chan<- string, a []string) {
	defer wg.Done()

	totalvowels := 0

	var iwg sync.WaitGroup
	//var mu sync.Mutex

	vowels := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		iwg.Add(1)

		go func(s string, iwg *sync.WaitGroup, ind int) {
			defer iwg.Done()
			//vowels := 0
			for i := 0; i < len(s); i++ {
				if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {

					vowels[ind]++

				}
			}
			// mu.Lock()
			// totalvowels += vowels
			// mu.Unlock()

		}(a[i], &iwg, i)
	}
	iwg.Wait()
	for _, count := range vowels {
		totalvowels += count
	}
	resp := fmt.Sprintf("%d total vowels:", totalvowels)
	fmt.Println(resp)
	c <- resp
}

// func LineCounter(s string, wg *sync.WaitGroup, c chan<- string) {
// 	defer wg.Done()
// 	lines := 0
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '.' {
// 			lines++
// 		}
// 	}
// 	resp := fmt.Sprintf("%d total lines :", lines)
// 	fmt.Println(resp)
// 	c <- resp
// }

func Wordfrequeny(s string, wg *sync.WaitGroup, c chan<- string, a []string) {
	defer wg.Done()

	var iwg sync.WaitGroup

	// totalwordcount := make(map[string]int, len(a))

	wordscount := make([]map[string]int, len(a))
	for i := range wordscount {
		wordscount[i] = make(map[string]int)
	}
	temp := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		iwg.Add(1)
		go func(s string, iwg *sync.WaitGroup, ind int) {
			defer iwg.Done()
			for j := 0; j < len(s); j++ {
				if s[j] == ' ' || j == len(s)-1 {
					if val, key := wordscount[ind][temp[ind]]; key {
						wordscount[ind][temp[ind]] = val + 1
					} else {
						wordscount[ind][temp[ind]] = 1
					}
					temp[ind] = ""
				} else {
					temp[ind] += string(s[j])
				}
			}

		}(a[i], &iwg, i)

	}
	iwg.Wait()
	// for index, val := range wordscount {
	// 	fmt.Println(index, val)
	// }

	// for k := 0; k < len(wordscount); k++ {

	// }
	//var mergemap map[string]int
	mergemap := make(map[string]int)
	for i := 0; i < len(wordscount); i++ {

		for key := range wordscount[i] {

			if _, ok := mergemap[key]; ok {

				mergemap[key] = mergemap[key] + wordscount[i][key]
			} else {
				mergemap[key] = wordscount[i][key]
			}

		}
	}
	//fmt.Println(mergemap)

	resp := fmt.Sprintf("%v word frequencies:", mergemap)
	fmt.Println(resp)
	c <- resp
}

func Wordcounter(s string, wg *sync.WaitGroup, c chan<- string, a []string) {

	defer wg.Done()

	totalwords := 0
	var iwg sync.WaitGroup
	//var mu sync.Mutex
	words := make([]int, len(a))
	//spaces := make([]int, len())
	for i := 0; i < len(a); i++ {
		iwg.Add(1)
		go func(s string, iwg *sync.WaitGroup, ind int) {
			defer iwg.Done()
			//spaces := 0
			//words := 0
			for j := 0; j < len(s); j++ {
				if s[j] == ' ' || j == len(s)-1 {
					words[ind]++
				}
			}

			// mu.Lock()
			// words = spaces + 1
			// totalwords += words
			// mu.Unlock()

		}(a[i], &iwg, i)
	}
	iwg.Wait()
	for _, count := range words {
		totalwords += count
	}

	resp := fmt.Sprintf("%d total words:", totalwords)
	fmt.Println(resp)
	c <- resp
}
