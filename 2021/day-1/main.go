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
	flag.StringVar(&file, "file", "", "file to read")
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
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] > input[i] {
			counter++
		}
	}
	fmt.Println(counter)
}
