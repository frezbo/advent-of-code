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
	var stage int
	flag.StringVar(&file, "file", "input.txt", "file to read")
	flag.IntVar(&stage, "stage", 1, "stage to run")
	flag.Parse()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	var input []int
	for _, s := range strings.Split(string(content), "\n") {
		if s != "" {
			i, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				return
			}

			input = append(input, int(i))
		}

	}

	counter := 0

	switch stage {
	case 1:

		for i := 0; i < len(input)-1; i++ {
			if input[i+1] > input[i] {
				counter++
			}
		}
	case 2:
		for i := 0; i < len(input)-3; i++ {
			if (input[i+1] + input[i+2] + input[i+3]) > (input[i] + input[i+1] + input[i+2]) {
				counter++
			}
		}
	default:
		fmt.Println("Invalid stage")
	}

	fmt.Println(counter)
}
