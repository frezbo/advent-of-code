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

	// horizonal pos, cost
	fuelCosts := make(map[int]int)
	for i := 1; i <= largest(horizonalPos); i++ {
		// for i := 2; i <= 2; i++ {
		cost := 0
		for _, j := range horizonalPos {
			delta := j - i
			if delta < 0 {
				delta *= -1
			}
			cost += delta
		}
		fuelCosts[i] = cost
	}
	smallest := fuelCosts[1]
	for _, v := range fuelCosts {
		if v < smallest {
			smallest = v
		}
	}
	fmt.Println(smallest)

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
