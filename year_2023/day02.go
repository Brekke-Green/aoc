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

	dataArr := strings.Split(string(data), "Game ")
	dataArr = dataArr[1:]


	realGames := 0
	dataMap := make(map[int][]string)

	for i, line := range dataArr {
		cleaned := strings.SplitAfter(line, ": ")
		games := strings.Split(cleaned[1], "; ")
		dataMap[i] = games 
	}


	for num, games := range dataMap {
		for _, game := range games {
			flag := true
			cube := strings.Split(game, ",")
			for _, c := range cube {
				die := strings.Split(c, " ")
				n, err := strconv.Atoi(die[0])
				if err != nil {
					os.Exit(1)
				}
				if cubes[die[1]] < n {
					flag = false
				}
			}
			if flag {
				realGames+=num
			}
		}
	}

	fmt.Println(realGames)
}
