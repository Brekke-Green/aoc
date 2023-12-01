package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	result := 0
	data, err := os.ReadFile("day01_input.txt")
	if err != nil {
		os.Exit(1)
	}

	nums := map[string]int {
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	dataArr := strings.Fields(string(data))
	for i := 0; i < len(dataArr); i++ {
		first := 0
		second := 0
		// find numbers in string
		// iterate over string and check if is in nums (from bottom up and top down)
		for j := 0; j < len(dataArr[i]); j++ {
			if num, found := nums[string(dataArr[i][j])]; found {
				first = num
				break
			}
		}

		for k := len(dataArr[i])-1; k >= 0; k-- {
			if num, found := nums[string(dataArr[i][k])]; found {
				second = num
				break
			}
		}
		// combine numbers and add to result
		result+=(first * 10 + second)
	}

	fmt.Println(result)
	duration := time.Since(start)
	fmt.Println(duration)
}
