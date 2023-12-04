package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	data, err := os.ReadFile("day01_input.txt")
	if err != nil {
		os.Exit(1)
	}

	nums := map[string]int {
		"0":		0,
		"1":		1,
		"2":		2,
		"3":		3,
		"4":		4,
		"5":		5,
		"6":		6,
		"7":		7,
		"8":		8,
		"9":		9,
		"zero":		0,
		"one":		1,
		"two":		2,
		"three":	3,
		"four":		4,
		"five":		5,
		"six":		6,
		"seven":	7,
		"eight":	8,
		"nine":		9,
	}

//	charStart := map[string][]string {
//		"z":	[]string{"zero"},
//		"o":	[]string{"zero", "one"},
//		"t":	[]string{"eight", "two", "three"},
//		"f":	[]string{"four", "five"},
//		"s":	[]string{"seven", "six"},
//		"e":	[]string{"eight", "three", "five", "one"},
//		"n":	[]string{"seven", "nine"},
//		"r":	[]string{"four"},
//		"x":	[]string{"six"},
//	}

	dataArr := strings.Fields(string(data))

	results := make([]int, 0)

	strChan := make(chan string)
	resChan := make(chan int)
	// errChan := make(chan error)
	doneChan := make(chan struct{})

	wg := sync.WaitGroup{}

	go func() {
		defer close(strChan)
		for _, s := range dataArr {
			strChan <- s
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done() 

			for s := range strChan {
				first := 0
				second := 0
				// find numbers in string
				// iterate over string and check if is in nums (from bottom up and top down)
				for j := 0; j < len(s); j++ {
					if num, found := nums[string(s[j])]; found {
						first = num
						break
					}
				}

				for k := len(s)-1; k >= 0; k-- {
					if num, found := nums[string(s[k])]; found {
						second = num
						break
					}
				}
				// combine numbers and add to result
				result := (first * 10 + second)
				// send result to resChan
				resChan <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	for {
		select {
		case result := <-resChan:
			results = append(results, result)
		case <-doneChan:
			sum := 0
			for _, num := range results {
				sum+=num
			}
			fmt.Println(sum)
			duration := time.Since(start)
			fmt.Println(duration)
			os.Exit(1)
		}
	}
}
