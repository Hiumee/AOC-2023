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

	for _, card := range data {
		parts := strings.Split(card, ":")
		numbers := parts[1]

		numbersParts := strings.Split(numbers, "|")
		winningNumbers, myNumbers := numbersParts[0], numbersParts[1]

		winningSet := make(map[int]struct{})

		winning := strings.Split(winningNumbers[1:len(winningNumbers)-1], " ")

		for _, win := range winning {
			n, err := strconv.Atoi(win)
			if err == nil {
				winningSet[n] = struct{}{}
			}
		}

		points := 0
		mine := strings.Split(myNumbers[1:], " ")
		for _, m := range mine {
			n, err := strconv.Atoi(m)
			if err == nil {
				if _, ok := winningSet[n]; ok {
					points *= 2
					if points == 0 {
						points = 1
					}
				}
			}
		}
		s += points
	}

	fmt.Println("Part 1:", s)

	// Part 2

	s = len(data)
	cardsWon := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		parts := strings.Split(data[i], ":")
		numbers := parts[1]

		numbersParts := strings.Split(numbers, "|")
		winningNumbers, myNumbers := numbersParts[0], numbersParts[1]

		winningSet := make(map[int]struct{})

		winning := strings.Split(winningNumbers[1:len(winningNumbers)-1], " ")

		for _, win := range winning {
			n, err := strconv.Atoi(win)
			if err == nil {
				winningSet[n] = struct{}{}
			}
		}

		points := 0
		mine := strings.Split(myNumbers[1:], " ")
		for _, m := range mine {
			n, err := strconv.Atoi(m)
			if err == nil {
				if _, ok := winningSet[n]; ok {
					points += 1
				}
			}
		}

		currentPoints := 0
		for j := 0; j < points && i+1+j < len(data); j++ {
			currentPoints += cardsWon[i+1+j] + 1
		}

		cardsWon[i] = currentPoints
		s += currentPoints
	}

	fmt.Println("Part 2:", s)
}
