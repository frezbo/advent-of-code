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
	currentPos := []int{0, 0}
	for _, s := range strings.Split(string(content), "\n") {
		if s != "" {
			commandInfo := strings.Split(s, " ")
			pos, err := strconv.Atoi(commandInfo[1])
			if err != nil {
				fmt.Println(err)
				return
			}
			switch stage {
			case 1:
				switch commandInfo[0] {
				case "forward":
					currentPos[0] += pos
				case "down":
					currentPos[1] += pos
				case "up":
					currentPos[1] -= pos
				default:
					fmt.Println("Unknown command")
					return
				}
			case 2:
			default:
				fmt.Println("Invalid stage")
			}
		}

	}

	fmt.Println(currentPos[0] * currentPos[1])
}
