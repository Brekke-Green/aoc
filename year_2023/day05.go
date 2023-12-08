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
	//fmt.Println(string(data))
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
	//fmt.Printf("%#v\n", seeds)


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
	
	for i, _ := range seeds {
		//fmt.Printf("INITIAL SEED => %v\n", seed)
		// run through mapping
		for j, mapping := range mappings {

			if (seeds[i] >= mapping[1] && seeds[i] <= (mapping[1]+mapping[2]-1)) {
				seeds[i] = seeds[i] + (mapping[0] - mapping[1])
				fmt.Printf("%v from %v\n", seeds[i], mappings[j])
			}

		}

	}
	fmt.Printf("%#v\n", seeds)
	lowest := seeds[0]
	for _, n := range seeds {
		if n < lowest {
			lowest = n
		}
	}
	fmt.Println(lowest)

	// find smallest seed and print

}
