package main

import (
	"fmt"
)

type Doc struct {
	time	[]int
	dist	[]int
}

func main() {
	//sample := Doc{time:	[]int{7, 15, 30}, dist: []int{9, 40, 200}}
	input := Doc{time:	[]int{59, 70, 78, 78}, dist: []int{430, 1218, 1213, 1276}}

	var results []int

	for i, t := range input.time {
		wins := 0
		for j := 1; j < t; j++ {
			d := j * (t - j)
			if d > input.dist[i] {
				wins+=1
			}
		}
		results = append(results, wins)
	}

	result := 1
	for _, num := range results {
		result*=num
	}
	fmt.Printf("Margin of error = %v\n", result)
}
