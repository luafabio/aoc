package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/samber/lo"
)

func main() {
	filePath := "C:\\Users\\Luis\\go\\aoc\\aoc.txt"

	b, _ := os.ReadFile(filePath)
	result := partone(b)
	fmt.Println(result)
}

func partone(b []byte) int {
	score := 0
	for _, line := range bytes.Split(b, []byte("\n")) {
		score += calculate(line)
	}
	return score
}

func calculate(input []byte) int {
	search := []byte("0123456789")
	result := make([]byte, 0, len(input))

	for i, _ := range input {
		if !unicode.IsDigit(rune(input[i])) {
			if i+2 <= len(input) {
				if input[i] == 'o' && input[i+1] == 'n' && input[i+2] == 'e' {
					input[i] = '1'
					input[i+1] = '1'
					input[i+2] = '1'
				}
				if input[i] == 't' && input[i+1] == 'w' && input[i+2] == 'o' {
					input[i] = '2'
					input[i+1] = '2'
					input[i+2] = '2'
				}
				if input[i] == 's' && input[i+1] == 'i' && input[i+2] == 'x' {
					input[i] = '6'
					input[i+1] = '6'
					input[i+2] = '6'
				}
			}

			if i+3 <= len(input) {
				if input[i] == 'f' && input[i+1] == 'o' && input[i+2] == 'u' && input[i+3] == 'r' {
					input[i] = '4'
					input[i+1] = '4'
					input[i+2] = '4'
					input[i+3] = '4'
				}
				if input[i] == 'f' && input[i+1] == 'i' && input[i+2] == 'v' && input[i+3] == 'e' {
					input[i] = '5'
					input[i+1] = '5'
					input[i+2] = '5'
					input[i+3] = '5'
				}
				if input[i] == 'n' && input[i+1] == 'i' && input[i+2] == 'n' && input[i+3] == 'e' {
					input[i] = '9'
					input[i+1] = '9'
					input[i+2] = '9'
					input[i+3] = '9'
				}
			}

			if i+4 <= len(input) {
				if input[i] == 't' && input[i+1] == 'h' && input[i+2] == 'r' && input[i+3] == 'e' && input[i+4] == 'e' {
					input[i] = '3'
					input[i+1] = '3'
					input[i+2] = '3'
					input[i+3] = '3'
					input[i+4] = '3'
				}
				if input[i] == 's' && input[i+1] == 'e' && input[i+2] == 'v' && input[i+3] == 'e' && input[i+4] == 'n' {
					input[i] = '7'
					input[i+1] = '7'
					input[i+2] = '7'
					input[i+3] = '7'
					input[i+4] = '7'
				}
				if input[i] == 'e' && input[i+1] == 'i' && input[i+2] == 'g' && input[i+3] == 'h' && input[i+4] == 't' {
					input[i] = '8'
					input[i+1] = '8'
					input[i+2] = '8'
					input[i+3] = '8'
					input[i+4] = '8'
				}
			}
		}
	}

	for _, c := range input {
		if lo.Contains(search, c) {
			result = append(result, c)
		}
	}

	number := string(result[0]) + string(result[len(result)-1])
	num, _ := strconv.Atoi(number)
	fmt.Printf("%q :%v \n", input, num)

	return num
}
