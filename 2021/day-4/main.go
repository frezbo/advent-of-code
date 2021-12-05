package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

	var randomInput []int

	for _, r := range strings.Split(inputData[0], ",") {
		num, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		randomInput = append(randomInput, num)
	}

	var boards [][][]map[int]bool

	for _, s := range strings.Split(string(content), "\n\n")[1:] {
		if s != "" {
			var board [][]map[int]bool
			for _, r := range strings.Split(s, "\n") {
				var row []map[int]bool
				for _, j := range strings.Split(r, " ") {
					if j == "" {
						continue
					}
					num, err := strconv.Atoi(j)
					if err != nil {
						log.Fatal(err)
						return
					}
					present := map[int]bool{num: false}
					row = append(row, present)
				}
				board = append(board, row)
			}
			boards = append(boards, board)
		}
	}
	if soln := firstWinner(boards, randomInput); soln != -1 {
		fmt.Println(soln)
	} else {
		fmt.Println("should never reach here")
	}
}

func firstWinner(boards [][][]map[int]bool, randomInput []int) int {
	for i := 0; i < len(randomInput); i++ {
		for b := 0; b < len(boards); b++ {
			for c := 0; c < len(boards[b]); c++ {

				for r := 0; r < len(boards[b][c]); r++ {
					for k, _ := range boards[b][c][r] {

						if k == randomInput[i] {
							boards[b][c][r][k] = true
						}

						if allMarked(boards[b]) {
							var sum int
							for c := 0; c < len(boards[b]); c++ {
								for r := 0; r < len(boards[b][c]); r++ {
									for k, v := range boards[b][c][r] {
										if !v {
											sum += k
										}
									}
								}
							}

							return sum * randomInput[i]
						}

					}

				}
			}
		}
	}
	return -1
}

func allMarked(board [][]map[int]bool) bool {
	// column wise
	for c := 0; c < len(board); c++ {
		counter := 0
		for r := 0; r < len(board[c]); r++ {
			for _, found := range board[c][r] {
				if found {
					counter++
				}
				if counter == 5 {
					return true
				}
			}
		}
	}

	// row wise
	for r := 0; r < len(board); r++ {
		counter := 0
		for c := 0; c < len(board[r]); c++ {
			for _, found := range board[c][r] {
				if found {
					counter++
				}
				if counter == 5 {
					return true
				}
			}
		}
	}
	return false
}
