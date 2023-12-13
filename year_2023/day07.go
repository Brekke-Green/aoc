package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Hand struct {
	cards	string
	bid		int
	counts	map[string]int
	score	[]int
	rank	int
}

func main() {
	//cards := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	cards := map[string]int{"2":1, "3":2, "4":3, "5":4, "6":5, "7":6, "8":7, "9":8, "A":13, "J":10, "K":12, "Q":11, "T":9}

	//data, err := os.ReadFile("day07_sample.txt")
	data, err := os.ReadFile("day07_input.txt")
	if err != nil {
		panic(err)
	}

	dataArr := strings.Split(string(data), "\n")
	dataArr = dataArr[:len(dataArr)-1]

	var hands []Hand

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
		count := make(map[string]int)
		for _, card := range hand.cards {
			count[string(card)]+=1
		}
		hand.counts = count

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
			hand.score = append(hand.score, cards[string(c)])
		}

		// append to array of Hands
		hands = append(hands, hand)
	}

	// sort arr of hands	
	var sortedHands []Hand
	
	for _, hand := range hands {
		sortedHands = append(sortedHands, hand)
	}

	slices.SortFunc(sortedHands, func(a, b Hand) int {
		if n := slices.Compare(a.score, b.score); n != 0 {
			return n
		}
		// If names are equal, order by age
		return 0
	})
	

	// tabulate score
	for i, _ := range sortedHands {
		sortedHands[i].rank = i+1 
	}

	result := 0
	for _, h := range sortedHands {
		result+=h.bid*h.rank
	}

	fmt.Printf("Result is -> %v", result)
}
