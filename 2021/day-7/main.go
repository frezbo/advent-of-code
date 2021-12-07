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

	inputData := strings.Split(string(content), "\n")[0]

	var horizonalPos []int
	for _, s := range strings.Split(inputData, ",") {
		if s != "" {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				return
			}
			horizonalPos = append(horizonalPos, num)
		}
	}

	fmt.Println(findSmallestCost(horizonalPos, false))
	fmt.Println(findSmallestCost(horizonalPos, true))
}

func largest(values []int) int {
	largest := 0
	for _, i := range values {
		if i > largest {
			largest = i
		}
	}
	return largest
}

func findSmallestCost(input []int, cumulative bool) int {
	fuelCosts := make(map[int]int)
	for i := 1; i <= largest(input); i++ {
		cost := 0
		for _, j := range input {
			delta := j - i
			if delta < 0 {
				delta *= -1
			}
			if cumulative {
				cost += (delta * (delta + 1)) / 2
			} else {
				cost += delta
			}
		}
		fuelCosts[i] = cost
	}
	smallest := fuelCosts[1]
	for _, v := range fuelCosts {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}
