package main

import (
	"fmt"
)

func main() {
	data := make([]string, 0)
	digit_words := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for {
		var input string
		n, _ := fmt.Scanln(&input)
		// if err != nil {
		// 	panic("An error occured")
		// }
		if n == 0 {
			break
		}
		data = append(data, input)
	}
	output := make([]int, 0)

	for _, code := range data {
		c := 0
	out:
		for ch := 0; ch < len(code); ch++ {
			char := code[ch]
			if char >= '0' && char <= '9' {
				c += int(char - '0')
				break
			}
			for index, digit_word := range digit_words {
				if code[ch:min(ch+len(digit_word), len(code))] == digit_word {
					c += index
					break out
				}
			}
		}
		c *= 10
	out2:
		for ch := len(code) - 1; ch >= 0; ch-- {
			char := code[ch]
			if char >= '0' && char <= '9' {
				c += int(char - '0')
				break
			}
			for index, digit_word := range digit_words {
				if code[ch:min(ch+len(digit_word), len(code))] == digit_word {
					c += index
					break out2
				}
			}

		}
		output = append(output, c)
	}

	sum := 0
	for _, o := range output {
		sum += o
	}

	fmt.Println(sum)
}
