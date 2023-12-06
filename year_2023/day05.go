package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day05_sample.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(data))
	dataArr := strings.Split(string(data), "\n")

	// handle seeds
	seedstrings := strings.Split(dataArr[0], " ")[1:]
	var seeds []int64
	for _, seed := range seedstrings {
		s, err := strconv.ParseInt(seed, 10, 64)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, s)
	}
	fmt.Printf("%#v\n", seeds)


	// handle parsing maps
	var mappingStrings []string
	for _, line := range dataArr {
		if line != "" && !strings.Contains(line, ":") {
			mappingStrings = append(mappingStrings, line)
		}
	}
	
	var mappings [][]int64

	for _, s := range mappingStrings {
		sArr := strings.Split(s, " ")
		var intsArr []int64
		for _, str := range sArr {
			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				panic(err)
			}
			intsArr = append(intsArr, num)
		}
		mappings = append(mappings, intsArr)
	}
	// run the seeds through mapping
	for i, seed := range seeds {
		fmt.Printf("INITIAL SEED => %v\n", seed)
		flag := false 
		if seed == 14 {
			flag = true
		}
		// run through mapping
		for _, mapping := range mappings {

			if flag {
				fmt.Printf("%v < %v < %v == %v\n", mapping[1], seeds[i], (mapping[1]+mapping[2]), (seeds[i] >= mapping[1] && seeds[i] < (mapping[1]+mapping[2])))
			}
			if (seeds[i] >= mapping[1] && seeds[i] < (mapping[1]+mapping[2])) {
				if flag {
					fmt.Printf("%v\n", seeds[i])
				}
				seeds[i] = mapping[0] + (seeds[i] - mapping[1])
				if flag {
					fmt.Printf("%v\n", seeds[i])
				}
			}
		}

	}
	fmt.Printf("%#v", seeds)

	// find smallest seed and print

}
