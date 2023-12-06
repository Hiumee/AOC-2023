package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := make([]string, 0)

	gaps := 8
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			input := scanner.Text()
			if len(input) == 0 {
				gaps--
				if gaps == 0 {
					break
				}
			}
			data = append(data, input)
		}
	}

	seedsList := strings.Split(data[0], ": ")[1]
	seedsStrings := strings.Split(seedsList, " ")
	seeds := make(map[int]int)

	for _, seed := range seedsStrings {
		s, err := strconv.Atoi(seed)
		if err != nil {
			panic("Not a number")
		}
		seeds[s] = -1
	}

	dataOffset := 0
	for ; data[dataOffset] != ""; dataOffset++ {
	}

	for i := 0; i < 7; i++ {
		dataOffset += 2

		for ; dataOffset < len(data) && data[dataOffset] != ""; dataOffset++ {
			numbersStrings := strings.Split(data[dataOffset], " ")
			numbers := make([]int, 0)
			for _, n := range numbersStrings {
				number, err := strconv.Atoi(n)
				if err != nil {
					panic("Not a number")
				}
				numbers = append(numbers, number)
			}

			startDest := numbers[0]
			startSrc := numbers[1]
			length := numbers[2]

			for k := range seeds {
				if k >= startSrc && k < startSrc+length {
					seeds[k] = k - startSrc + startDest
				}
			}
		}

		mapping := make(map[int]int)
		for k, v := range seeds {
			if v == -1 {
				mapping[k] = -1
			} else {
				mapping[v] = -1
			}
		}
		seeds = mapping
	}

	mini := -1

	for k := range seeds {
		if mini == -1 || k < mini {
			mini = k
		}
	}

	fmt.Println(mini)
}
