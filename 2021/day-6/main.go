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
	var days int
	flag.StringVar(&file, "file", "input.txt", "file to read")
	flag.IntVar(&days, "days", 80, "number of days passed")
	flag.Parse()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	var input []int

	for _, s := range strings.Split(string(content), "\n") {
		if s != "" {
			for _, t := range strings.Split(s, ",") {
				num, err := strconv.Atoi(t)
				if err != nil {
					fmt.Println(num)
					return
				}
				input = append(input, num)
			}
		}
	}

	prev := input
	var result [][]int
	for i := 0; i < days; i++ {
		var k []int
		for _, val := range prev {
			lv := val
			lv--
			k = append(k, lv)
		}

		for i, v := range k {
			if v < 0 {
				k[i] = v + 7
				k = append(k, 8)
			}

		}
		prev = k
		result = append(result, k)
	}

	// fmt.Println(input)
	fmt.Println(len(result[len(result)-1]))
}
