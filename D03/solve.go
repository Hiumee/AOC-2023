package main

import (
	"bufio"
	"fmt"
	"os"
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

	for i, line := range data {
		partNumber := false
		number := 0
		for j, chr := range line {
			if chr >= '0' && chr <= '9' {
				if number == 0 {
					if j > 0 {
						if i > 0 && (data[i-1][j-1] < '0' || data[i-1][j-1] > '9') && data[i-1][j-1] != '.' {
							partNumber = true
						}
						if i < len(data)-1 && (data[i+1][j-1] < '0' || data[i+1][j-1] > '9') && data[i+1][j-1] != '.' {
							partNumber = true
						}
						if (data[i][j-1] < '0' || data[i][j-1] > '9') && data[i][j-1] != '.' {
							partNumber = true
						}
					}
				}
				number *= 10
				number += int(chr - '0')

				if i > 0 && (data[i-1][j] < '0' || data[i-1][j] > '9') && data[i-1][j] != '.' {
					partNumber = true
				}
				if i < len(data)-1 && (data[i+1][j] < '0' || data[i+1][j] > '9') && data[i+1][j] != '.' {
					partNumber = true
				}
			}
			if (chr < '0' || chr > '9') || j == len(line)-1 {
				if number != 0 && j < len(line)-1 {
					if i > 0 && (data[i-1][j] < '0' || data[i-1][j] > '9') && data[i-1][j] != '.' {
						partNumber = true
					}
					if i < len(data)-1 && (data[i+1][j] < '0' || data[i+1][j] > '9') && data[i+1][j] != '.' {
						partNumber = true
					}
					if (data[i][j] < '0' || data[i][j] > '9') && data[i][j] != '.' {
						partNumber = true
					}
				}
				if partNumber {
					s += number
				}
				number = 0
				partNumber = false
			}
		}
	}

	fmt.Println(s)
}
