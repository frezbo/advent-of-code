package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Coordinates struct {
	Verticals, Horizontals []Coordinate
}

type CoordinatePair struct {
	Start, End Coordinate
}

type Coordinate struct {
	X, Y int
}

func main() {
	var file string
	flag.StringVar(&file, "file", "input.txt", "file to read")
	flag.Parse()

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	inputData := strings.Split(string(content), "\n")

	var coordinates Coordinates

	for _, line := range inputData {
		if line != "" {
			coordinateData := strings.Split(line, "->")
			x1y1 := strings.Split(coordinateData[0], ",")
			x2y2 := strings.Split(coordinateData[len(coordinateData)-1], ",")
			x1, err := strconv.Atoi(strings.Trim(x1y1[0], " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			y1, err := strconv.Atoi(strings.Trim(x1y1[1], " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			x2, err := strconv.Atoi(strings.Trim(x2y2[0], " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			y2, err := strconv.Atoi(strings.Trim(x2y2[1], " "))
			if err != nil {
				fmt.Println(err)
				return
			}
			if x1 == x2 {
				var coordinatePair CoordinatePair
				if y1 < y2 {
					coordinatePair.Start = Coordinate{
						X: x1,
						Y: y1,
					}
					coordinatePair.End = Coordinate{
						X: x2,
						Y: y2,
					}
				} else {
					coordinatePair.End = Coordinate{
						X: x1,
						Y: y1,
					}
					coordinatePair.Start = Coordinate{
						X: x2,
						Y: y2,
					}
				}
				for i := coordinatePair.Start.Y; i <= coordinatePair.End.Y; i++ {
					coordinates.Verticals = append(coordinates.Verticals, Coordinate{
						X: coordinatePair.Start.X,
						Y: i,
					})
				}

			}

			if y1 == y2 {
				var coordinatePair CoordinatePair
				if x1 < x2 {
					coordinatePair.Start = Coordinate{
						X: x1,
						Y: y1,
					}
					coordinatePair.End = Coordinate{
						X: x2,
						Y: y2,
					}
				} else {
					coordinatePair.End = Coordinate{
						X: x1,
						Y: y1,
					}
					coordinatePair.Start = Coordinate{
						X: x2,
						Y: y2,
					}
				}
				for i := coordinatePair.Start.X; i <= coordinatePair.End.X; i++ {
					coordinates.Horizontals = append(coordinates.Horizontals, Coordinate{
						X: i,
						Y: coordinatePair.Start.Y,
					})
				}
			}
		}
	}

	seen := map[Coordinate]int{}
	for _, c := range coordinates.Horizontals {
		seen[c]++
	}
	for _, c := range coordinates.Verticals {
		seen[c]++
	}

	var count int
	for _, v := range seen {
		if v >= 2 {
			count++
		}
	}
	fmt.Println(count)

}
