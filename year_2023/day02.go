package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day02_input.txt")
	if err != nil {
		os.Exit(1)
	}

	cubes := map[string]int{
		"red":		12,
		"green":	13,
		"blue":		14,
	}

	data_string := strings.ReplaceAll(string(data), "\n", "")
	dataArr := strings.Split(data_string, "Game ")
	dataArr = dataArr[1:]


	realGames := 0
	badGamesMap := make(map[int]bool)
	dataMap := make(map[int][]string)

	for i, line := range dataArr {
		cleaned := strings.SplitAfter(line, ": ")
		games := strings.Split(cleaned[1], "; ")
		dataMap[i+1] = games 
	}
	


	for num, games := range dataMap {
		for _, game := range games {
			cube := strings.Split(game, ", ")
			for _, c := range cube {
				die := strings.Split(c, " ")
				die[1] = strings.ReplaceAll(die[1], "\n", "")
				n, err := strconv.Atoi(die[0])
				if err != nil {
					panic(err)
				}
				if cubes[die[1]] < n {
					badGamesMap[num] = true
				}
			}
		}
	}

	for i := 1; i <= 100; i++ {
		if _, found := badGamesMap[i]; found == false {
			realGames+=i
		}
	}

	fmt.Println(realGames)
}
