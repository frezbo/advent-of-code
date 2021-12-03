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

	valuesToConsider := strings.Split(string(content), "\n")

	// should be a better way than checking the input
	bitLength := len(valuesToConsider[0])

	commonBits := make([]map[int]int, bitLength)
	for i := range commonBits {
		commonBits[i] = make(map[int]int)
	}

	for _, s := range valuesToConsider {
		if s != "" {
			for i, b := range s {

				bit, err := strconv.Atoi(string(b))
				if err != nil {
					fmt.Println(err)
					return
				}

				if bit == 1 {
					commonBits[i][1]++
				} else {
					commonBits[i][0]++
				}
			}
		}

	}
	gammaBitArray := []string{}
	for _, m := range commonBits {
		if m[0] > m[1] {
			gammaBitArray = append(gammaBitArray, "0")
		} else {
			gammaBitArray = append(gammaBitArray, "1")
		}
	}
	i, err := strconv.ParseInt(strings.Join(gammaBitArray, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	epsilonBitArray := []string{}
	for _, m := range commonBits {
		if m[0] < m[1] {
			epsilonBitArray = append(epsilonBitArray, "0")
		} else {
			epsilonBitArray = append(epsilonBitArray, "1")
		}
	}
	j, err := strconv.ParseInt(strings.Join(epsilonBitArray, ""), 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i * j)

	oxygenBinary := getOxygenOrCO2Rating(valuesToConsider, 0, false)[0]
	oxygen, err := strconv.ParseInt(oxygenBinary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	co2Binary := getOxygenOrCO2Rating(valuesToConsider, 0, true)[0]
	co2, err := strconv.ParseInt(co2Binary, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(co2 * oxygen)
}

func getOxygenOrCO2Rating(values []string, pos int, co2 bool) []string {
	if len(values) == 1 {
		return values
	}
	var newValues []string
	count := map[int]int{}
	for _, s := range values {
		if s != "" {
			firstBit := string(s[pos])
			bit, err := strconv.Atoi(firstBit)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			count[bit]++
		}
	}
	var greatest int
	if co2 {
		if count[0] <= count[1] {
			greatest = 0
		} else {
			greatest = 1
		}
	} else {
		if count[1] >= count[0] {
			greatest = 1
		} else {
			greatest = 0
		}
	}
	for _, s := range values {
		if s != "" {
			firstBit := string(s[pos])
			bit, err := strconv.Atoi(firstBit)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			if bit == greatest {
				newValues = append(newValues, s)
			}
		}
	}

	pos++

	newValues = getOxygenOrCO2Rating(newValues, pos, co2)
	return newValues
}
