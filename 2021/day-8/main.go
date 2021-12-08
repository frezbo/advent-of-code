package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	segmentMap := map[int][]string{
		0: {"a", "b"},
		1: {"c", "f"},
		2: {"a", "c", "d", "e", "g"},
		3: {"a", "c", "d", "f", "g"},
		4: {"b", "c", "d", "f"},
		5: {"a", "b", "d", "f", "g"},
		6: {"a", "b", "d", "e", "f", "g"},
		7: {"a", "c", "f"},
		8: {"a", "b", "c", "d", "e", "f", "g"},
		9: {"a", "b", "c", "d", "f", "g"},
	}

	input := make(map[int]int)
	for _, s := range strings.Split(string(content), "\n") {
		if s != "" {
			for _, d := range strings.Split(strings.Split(s, " | ")[1], " ") {
				for s := 0; s <= 9; s++ {
					if len(segmentMap[s]) == len(d) {
						input[s]++
					}
				}
			}
		}

	}

	fmt.Println(input[1] + input[4] + input[7] + input[8])

}
