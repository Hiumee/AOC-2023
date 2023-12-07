package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Round struct {
	Hand  string
	Bid   int
	Score int
}

func GetScore(hand string) int {
	// 5 of a kind
	if hand[0] == hand[1] && hand[0] == hand[2] && hand[0] == hand[3] && hand[0] == hand[4] {
		return 7
	}

	numbers := make(map[rune]int)

	for _, b := range hand {
		if _, ok := numbers[b]; !ok {
			numbers[b] = 0
		}
		numbers[b]++
	}

	// 4 of a kind
	for _, v := range numbers {
		if v == 4 {
			return 6
		}
	}

	// full house
	if len(numbers) == 2 {
		return 5
	}

	// 3 of a kind
	for _, v := range numbers {
		if v == 3 {
			return 4
		}
	}

	// 2 pair
	if len(numbers) == 3 {
		return 3
	}

	// 1 pair
	if len(numbers) == 4 {
		return 2
	}

	// high card
	return 1
}

func GetScoreP2(hand string) int {
	numbers := make(map[rune]int)

	for _, b := range hand {
		if _, ok := numbers[b]; !ok {
			numbers[b] = 0
		}
		numbers[b]++
	}

	Js := numbers['J']
	delete(numbers, 'J')

	// 5 of a kind
	if len(numbers) <= 1 {
		return 7
	}

	// 4 of a kind
	for _, v := range numbers {
		if v+Js == 4 {
			return 6
		}
	}

	// full house
	if len(numbers) == 2 {
		return 5
	}

	// 3 of a kind
	for _, v := range numbers {
		if v+Js == 3 {
			return 4
		}
	}

	// 2 pair
	if len(numbers) == 3 {
		return 3
	}

	// 1 pair
	if len(numbers) == 4 {
		return 2
	}

	// high card
	return 1
}

func main() {
	cardScore := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	cardScoreP2 := map[byte]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 1,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

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

	rounds := make([]Round, 0)

	for _, line := range data {
		parts := strings.Split(line, " ")
		hand, bid := parts[0], parts[1]

		n, err := strconv.Atoi(bid)
		if err != nil {
			panic("Not a number")
		}

		rounds = append(rounds, Round{hand, n, GetScore(hand)})
	}

	sort.Slice(rounds, func(i, j int) bool {
		if rounds[i].Score != rounds[j].Score {
			return rounds[i].Score < rounds[j].Score
		}

		var k int
		for k = 0; k < 5; k++ {
			if rounds[i].Hand[k] != rounds[j].Hand[k] {
				return cardScore[rounds[i].Hand[k]] < cardScore[rounds[j].Hand[k]]
			}
		}
		return true
	})

	s := 0
	for rank, round := range rounds {
		s += (rank + 1) * round.Bid
	}

	fmt.Println("Part 1:", s)

	rounds = make([]Round, 0)

	for _, line := range data {
		parts := strings.Split(line, " ")
		hand, bid := parts[0], parts[1]

		n, err := strconv.Atoi(bid)
		if err != nil {
			panic("Not a number")
		}

		rounds = append(rounds, Round{hand, n, GetScoreP2(hand)})
	}

	sort.Slice(rounds, func(i, j int) bool {
		if rounds[i].Score != rounds[j].Score {
			return rounds[i].Score < rounds[j].Score
		}

		var k int
		for k = 0; k < 5; k++ {
			if rounds[i].Hand[k] != rounds[j].Hand[k] {
				return cardScoreP2[rounds[i].Hand[k]] < cardScoreP2[rounds[j].Hand[k]]
			}
		}
		return true
	})

	s = 0
	for rank, round := range rounds {
		s += (rank + 1) * round.Bid
	}

	fmt.Println("Part 2:", s)
}
