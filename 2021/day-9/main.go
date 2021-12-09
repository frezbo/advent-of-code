package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "input.txt", "file to read")
	flag.Parse()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	inputData := strings.Split(string(content), "\n")
	heightMap := make([][]int, 1)

	heightMap[0] = pad(heightMap[0], len(inputData[0])+2)

	for _, s := range inputData {
		if s != "" {
			height := make([]int, 1)
			height = pad(height, 1)
			for _, h := range strings.Split(s, "") {
				num, err := strconv.Atoi(h)
				if err != nil {
					fmt.Println(err)
					return
				}
				height = append(height, num)
			}
			height = pad(height, 1)
			heightMap = append(heightMap, height)
		}

	}
	heightMap = append(heightMap, pad([]int{}, len(inputData[0])+2))

	var count int
	for c := 1; c < len(heightMap)-1; c++ {
		for r := 1; r < len(heightMap[c])-1; r++ {
			if heightMap[c][r] < heightMap[c][r+1] && heightMap[c][r] < heightMap[c][r-1] && heightMap[c][r] < heightMap[c+1][r] && heightMap[c][r] < heightMap[c-1][r] {
				count += heightMap[c][r] + 1
			}
		}
	}

	fmt.Println(count)
}

func pad(input []int, count int) []int {
	for i := 0; i < count; i++ {
		input = append(input, 10)
	}
	return input
}
