package main

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	val		string
	left	*Node
	right	*Node
}

func main() {
	//data, err := os.ReadFile("day08_sample.txt")
	data, err := os.ReadFile("day08_input.txt")
	if err != nil {
		os.Exit(1)
	}

	dataArr := strings.Split(string(data), "\n")
	dataArr = dataArr[:len(dataArr)-1]

	instructions := strings.Split(dataArr[0], "")
	nodes := dataArr[2:]
	// create map from data values
	graph := make(map[string][]string)
	for _, row := range nodes {
		graph[row[:3]] = []string{row[7:10], row[12:15]}
	}

	// follow steps to get to end and if resulting node isn't at ZZZ repeat
	current := "AAA"
	steps := 0
	
	for current != "ZZZ" {
		for _, step := range instructions {
			if step == "L" {
				current = graph[current][0]
				steps+=1
			} else {
				current = graph[current][1]
				steps+=1
			}
		}
	}
	fmt.Printf("It took %v steps", steps)
}
