package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNumbersP1(line string) []int {
	numbers := strings.Split(strings.Split(line, ":")[1], " ")
	justNumbers := make([]int, 0)

	for _, k := range numbers {
		n, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		justNumbers = append(justNumbers, n)
	}

	return justNumbers
}

func getNumbersP2(line string) []int {
	numbers := strings.Split(strings.Split(line, ":")[1], " ")
	numberString := ""

	for _, k := range numbers {
		if k != "" {
			numberString += k
		}
	}

	n, err := strconv.Atoi(numberString)
	if err != nil {
		panic("Not a number")
	}

	return []int{n}
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

	timeStrings, distanceStrings := data[0], data[1]

	times := getNumbersP1(timeStrings)
	distances := getNumbersP1(distanceStrings)
	options := 1

	for i := 0; i < len(times); i++ {
		time := float64(times[i])
		distance := float64(distances[i])

		delta := math.Sqrt(time*time - 4*distance)
		minimum := int(math.Ceil((time - delta) / 2))
		maximum := int(math.Floor((time + delta) / 2))

		if maximum*(times[i]-maximum) <= distances[i] {
			maximum--
		}
		if minimum*(times[i]-minimum) <= distances[i] {
			minimum++
		}

		options *= maximum - minimum + 1
	}

	fmt.Println("Part 1:", options)

	times = getNumbersP2(timeStrings)
	distances = getNumbersP2(distanceStrings)
	options = 1

	for i := 0; i < len(times); i++ {
		time := float64(times[i])
		distance := float64(distances[i])

		delta := math.Sqrt(time*time - 4*distance)
		minimum := int(math.Ceil((time - delta) / 2))
		maximum := int(math.Floor((time + delta) / 2))

		if maximum*(times[i]-maximum) <= distances[i] {
			maximum--
		}
		if minimum*(times[i]-minimum) <= distances[i] {
			minimum++
		}

		options *= maximum - minimum + 1
	}

	fmt.Println("Part 2:", options)
}
