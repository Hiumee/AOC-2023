package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNumberAtPosition(data []string, i int, j int) (int, int, int) {
	if i < 0 || j < 0 || i >= len(data) || j >= len(data[i]) || !(data[i][j] >= '0' && data[i][j] <= '9') {
		return 0, 0, 0
	}

	for ; j >= 0 && data[i][j] >= '0' && data[i][j] <= '9'; j-- {
	}
	j++

	start := j
	number := 0

	for ; j < len(data[i]) && data[i][j] >= '0' && data[i][j] <= '9'; j++ {
		number *= 10
		number += int(data[i][j] - '0')
	}

	j--

	return number, start, j
}

func main() {
	data := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			input := scanner.Text()
			if len(input) == 0 {
				break
			}
			data = append(data, input)
		}
	}

	s := 0

	for i, line := range data {
		for j, chr := range line {
			if chr == '*' {
				numbers := make([]int, 0)

				if j > 0 {
					number, _, _ := getNumberAtPosition(data, i, j-1)
					numbers = append(numbers, number)
				}
				if j < len(line)-1 {
					number, _, _ := getNumberAtPosition(data, i, j+1)
					numbers = append(numbers, number)
				}

				top1, start1, _ := getNumberAtPosition(data, i-1, j-1)
				top2, start2, _ := getNumberAtPosition(data, i-1, j)
				top3, start3, _ := getNumberAtPosition(data, i-1, j+1)

				numbers = append(numbers, top1)

				if start1 != start2 {
					numbers = append(numbers, top2)
				}

				if start2 != start3 {
					numbers = append(numbers, top3)
				}

				bot1, start1, _ := getNumberAtPosition(data, i+1, j-1)
				bot2, start2, _ := getNumberAtPosition(data, i+1, j)
				bot3, start3, _ := getNumberAtPosition(data, i+1, j+1)

				numbers = append(numbers, bot1)

				if start1 != start2 {
					numbers = append(numbers, bot2)
				}

				if start2 != start3 {
					numbers = append(numbers, bot3)
				}

				nc := 0
				p := 1

				for _, n := range numbers {
					if n != 0 {
						nc++
						p *= n
					}
				}

				if nc == 2 {
					s += p
				}
			}
		}
	}

	fmt.Println(s)
}
