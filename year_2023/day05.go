package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	dest	int64
	src		int64
	r		int64
}

func main() {
	data, err := os.ReadFile("day05_input.txt")
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
	var mappingStrings [][]string
	var mapping []string
	for _, line := range dataArr {
		if strings.Contains(line, ":") {
			mapping = []string{}
		} else if line == "" {
			mappingStrings = append(mappingStrings, mapping)
		} else {
			mapping = append(mapping, line)
		}
	}
	
	var mappings [][]Mapping

	for j, s := range mappingStrings {
		if len(s) == 0 || j == len(mappingStrings) - 1 {
			continue
		}
		m := []Mapping{}
		for _, l := range s {
			sArr := strings.Split(l, " ")
			var row Mapping
			for i, str := range sArr {
				num, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					panic(err)
				}
				if i == 0 {row.dest = num}
				if i == 1 {row.src = num}
				if i == 2 {row.r = num}
			}
			m = append(m, row)
		}

		mappings = append(mappings, m)
	}
	//run the seeds through mapping
	
	for i, _ := range seeds {
	//	 run through mapping
		for _, group := range mappings {
			for _, mapping := range group {
				if (seeds[i] >= mapping.src && seeds[i] <= (mapping.src+mapping.r-1)) {
					seeds[i] = seeds[i] + (mapping.dest - mapping.src)
					break
				}
			}
		}
	}
	fmt.Printf("%#v\n", mappings)
	lowest := seeds[0]
	for _, n := range seeds {
		if n < lowest {
			lowest = n
		}
	}
	
	fmt.Printf("%#v\n", seeds)
	fmt.Println(lowest)

	// find smallest seed and print

}
