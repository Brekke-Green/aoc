package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"

	"golang.org/x/exp/slices"
)

func main() {
    data, err := os.ReadFile("day09_sample.txt")
    if err != nil {
        os.Exit(1)
    }

    dataArr := strings.Split(string(data), "\n")
    dataArr = dataArr[:len(dataArr)-1]

    var dataIntArr [][]int

    for _, row := range dataArr {
        var intsRow []int
        numsRow := strings.Split(row, " ")
        for _, digit := range numsRow {
            num, err := strconv.Atoi(digit)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            intsRow = append(intsRow, num)
        }
        dataIntArr = append(dataIntArr, intsRow)
    }
    
    // create next digit
	for i, row := range dataIntArr {
		dataIntArr[i] = nextDigit(&row)
	}

    // sum all final elements 

}

func createDifferences() {

}

func nextDigit(nums *[]int) []int {
	var stack [][]int
	d := *nums
	// find difference for last digit
	stack = append(stack, d)



	return d
}
