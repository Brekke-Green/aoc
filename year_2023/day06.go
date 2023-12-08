package main

import (
	"fmt"
)

type Doc struct {
	time	[]int64
	dist	[]int64
}

func main() {
	//sample := Doc{time:	[]int64{71530}, dist: []int64{940200}}
	input := Doc{time: []int64{59707878}, dist: []int64{430121812131276}}

	var results []int64

	for i, t := range input.time {
		var wins int64
		wins = 0
		var j int64
		for j = 1; j < t; j++ {
			d := j * (t - j)
			if d > input.dist[i] {
				wins+=1
			}
		}
		results = append(results, wins)
	}

	//result := 1
	//for _, num := range results {
	//	result*=num
	//}
	fmt.Printf("Margin of error = %v\n", results)
}
