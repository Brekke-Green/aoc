package main

import (
	"fmt"
	"os"
	"strings"
)

type Cord struct {
	x	int
	y	int
}

type Number struct {
	start	Cord
	length	int
}

func main() {

	numeric := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	data, err := os.ReadFile("day03_input.txt")
	if err != nil {
		os.Exit(1)
	}
	dataString := string(data)
	dataArr := strings.Split(dataString, "\n")
	dataArr = dataArr[:len(dataArr)-1]
	var data2D [][]string 

	for i := 0; i < len(dataArr); i++ {
		data2D = append(data2D, []string{""})
		for j := 0; j < len(dataArr[0]); j++ {
			data2D[i] = strings.Split(dataArr[i], "")
		}
	}

	var nums []Number

	yCord := 0
	xCord := 0

	for yCord < len(data2D)  {
		xCord = 0
		for xCord < len(data2D[0]) {
			xStart := xCord

			char := data2D[yCord][xCord]

			_, found := numeric[char] 
			s := found
			if s {
				var num Number
				num.start = Cord{xCord, yCord}
				// get length
				stillNum := true 
				length := 1
				for stillNum {
					stillNum = false
					var next string
					if xCord > len(data2D[0]) - 1 {
						next = "."
					} else {
						next = data2D[yCord][xCord]
					}
					_, found := numeric[next]
					if found {
						length+=1
						xCord+=1
						stillNum = true
					} else {
						num.length+=length-1
						xCord+=1
					}
				}
				// append Number to nums list
				nums = append(nums, num)
			}

			if xCord == xStart {
				xCord+=1
			}
			if xCord >= len(data2D) {
				yCord+=1
			}
		}
	}

	for _, num := range nums {
		fmt.Println(num)
	}
}

//func findNum(grid [][]string, x int, y int) [][]int {
//	var cords [][]int
//	cords = append(cords, []int{})
//	cords = append(cords, []int{})
//	return cords
//}

//func getNeighbors(xs []int, ys []int) {
//	
//}
