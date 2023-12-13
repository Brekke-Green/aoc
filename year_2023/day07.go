package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards	string
	bid		int
	counts	map[string]int
	score	[]int

}

func main() {
	//cards := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	cards := map[string]int{"2":1, "3":2, "4":3, "5":4, "6":5, "7":6, "8":7, "9":8, "A":13, "J":10, "K":12, "Q":11, "T":9}

	data, err := os.ReadFile("day07_sample.txt")
	//data, err := os.ReadFile("day07_input.txt")
	if err != nil {
		panic(err)
	}

	dataArr := strings.Split(string(data), "\n")
	dataArr = dataArr[:len(dataArr)-1]

	// iterate through each row and create Hand struct
	for _, row := range dataArr {
		arr := strings.Split(row, " ")
		hand := Hand{}
		hand.cards = arr[0]
		bid, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(err)
		}
		hand.bid = bid

		// count cards
		var count map[rune]int
		for _, card := range hand.cards {
			count[card]+=1
		}

		// score cards
		var counts []int
		for _, v := range count {
			counts = append(counts, v)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(counts)))

		if counts[0] == 5 {
			hand.score = append(hand.score, 6)
		} else if counts[0] == 4 {
			hand.score = append(hand.score, 5)
		} else if counts[0] == 1 {
			hand.score = append(hand.score, 0)
		} else if counts[0] == 3 && counts[1] == 2 {
			hand.score = append(hand.score, 4)
		} else if counts[0] == 3 && counts[1] == 1 {
			hand.score = append(hand.score, 3)
		} else if counts[0] == 2 && counts[1] == 2 {
			hand.score = append(hand.score, 2)
		} else if counts[0] == 2 && counts[1] == 1 {
			hand.score = append(hand.score, 1)
		}

		for _, c := range hand.cards {
			hand.score
		}


		// append to array of Hands
		
	}

	// sort arr of hands

	// tabulate score

	
	fmt.Printf("%v\n", cards)
	fmt.Printf("%#v\n", dataArr)
}
