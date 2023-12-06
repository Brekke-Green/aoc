package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day05_sample.txt")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(data))
	dataArr := strings.Split(string(data), "\n")
	fmt.Printf("%#v", dataArr)
}
