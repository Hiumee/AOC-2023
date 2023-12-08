package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Lr struct {
	Left  string
	Right string
}

func Gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Lcm(numbers []int) int {
	if len(numbers) < 2 {
		panic("Not enough numbers")
	}
	result := numbers[0] * numbers[1] / Gcd(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		result = result * numbers[i] / Gcd(result, numbers[i])
	}

	return result
}

func main() {
	data := make([]string, 0)

	gaps := 2
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

	steps := data[0]

	nodes := data[2:]

	graph := make(map[string]Lr)

	for _, node := range nodes {
		parts := strings.Split(node, " = ")
		name, sides := parts[0], parts[1]

		sidesStrings := strings.Split(sides[1:len(sides)-1], ", ")
		left, right := sidesStrings[0], sidesStrings[1]

		graph[name] = Lr{left, right}
	}

	if _, ok := graph["AAA"]; ok {
		stepNumber := 0
		current := "AAA"
		for current != "ZZZ" {
			for _, step := range steps {
				switch step {
				case 'L':
					current = graph[current].Left
				case 'R':
					current = graph[current].Right
				}
			}
			stepNumber += len(steps)
		}

		fmt.Println("Part 1:", stepNumber)
	}

	stepNumbers := make([]int, 0)
	for k := range graph {
		if k[2] == 'A' {
			current := k

			stepNumber := 0

			for current[2] != 'Z' {
				for _, step := range steps {
					switch step {
					case 'L':
						current = graph[current].Left
					case 'R':
						current = graph[current].Right
					}
				}
				stepNumber += len(steps)
			}

			stepNumbers = append(stepNumbers, stepNumber)
		}
	}

	result := Lcm(stepNumbers)

	fmt.Println("Step 2:", result)
}
