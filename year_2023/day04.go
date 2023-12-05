package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winners		[]int
	players		[]int
	score		int
}

func main() {
	data, err := os.ReadFile("day04_input.txt")	
	if err != nil {
		os.Exit(1)
	}

	dataArr := strings.Split(string(data), "\n")
	var data2D [][]string
	
	for i := 0; i < len(dataArr); i++ {
		data2D = append(data2D, strings.Split(dataArr[i], ": "))
	}
	data2D = data2D[:len(data2D)-1]
	
	var cards [][]string

	for i := 0; i < len(data2D); i++ {
		cards = append(cards, strings.Split(data2D[i][1], " | "))
	}

	var cardsList []Card
	for i := 0; i < len(cards); i++ {
		c := Card{}
		winners := strings.Split(cards[i][0], " ")
		players := strings.Split(cards[i][1], " ")
		winnersInt := []int{}
		playersInt := []int{}	

		for i, s := range winners {
			if s == "" {
				continue
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(i)
				panic(err)
			}
			winnersInt = append(winnersInt, num)
		}
		for i, s := range players {
			if s == "" {
				continue
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(i)
				panic(err)
			}
			playersInt = append(playersInt, num)
		}
		c.winners = winnersInt
		c.players = playersInt 
		cardsList = append(cardsList, c)
	}

	// SCORE THE CARD GAMES NOW

	fmt.Printf("%#v", cardsList)
}
