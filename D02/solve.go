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

	output := make([]int, 0)
	powers := make([]int, 0)

	global_r := 12
	global_g := 13
	global_b := 14

	for _, line := range data {
		parts := strings.Split(line, ":")
		game, cubes := parts[0], parts[1]

		game = game[5:]
		gameNumber, err := strconv.Atoi(game)
		if err != nil {
			panic("Can't convert to number")
		}

		cubes_list := strings.Split(cubes, ";")

		valid := true

		max_r := 0
		max_g := 0
		max_b := 0

		for _, cube_list := range cubes_list {
			parts := strings.Split(cube_list, ",")
			r := 0
			g := 0
			b := 0

			for _, part := range parts {
				number_color := strings.Split(part[1:], " ")
				numberA, color := number_color[0], number_color[1]

				number, err := strconv.Atoi(numberA)

				if err != nil {
					panic("Can't convert to number")
				}

				switch color {
				case "red":
					r += number
				case "green":
					g += number
				case "blue":
					b += number
				}
			}
			if r > max_r {
				max_r = r
			}
			if g > max_g {
				max_g = g
			}
			if b > max_b {
				max_b = b
			}
			if r > global_r || g > global_g || b > global_b {
				valid = false
			}
		}

		powers = append(powers, max_r*max_g*max_b)

		if valid {
			output = append(output, gameNumber)
		}

	}

	s := 0
	for _, n := range output {
		s += n
	}

	fmt.Println(s)

	s = 0
	for _, n := range powers {
		s += n
	}
	fmt.Println(s)
}
