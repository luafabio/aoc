package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := "C:\\Users\\Luis\\go\\aoc\\dayTwo\\aoc.txt"

	b, _ := os.ReadFile(filePath)
	result := partTwo(b)
	fmt.Println(result)
}

func partTwo(b []byte) int {
	score := 0

	for i, line := range bytes.Split(b, []byte("\n")) {
		game := bytes.Split(line, []byte(": "))
		turns := bytes.Split(game[1], []byte("; "))

		m := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, turn := range turns {

			colors := bytes.Split(turn, []byte(", "))

			for _, color := range colors {
				tmp := bytes.Split(color, []byte(" "))

				score, err := strconv.Atoi(string(tmp[0]))
				if err != nil {
					panic(err)
				}

				m[string(tmp[1])] += score
			}
		}

		result := Game{
			index: i,
			turn: Turn{
				red:   m["red"],
				blue:  m["blue"],
				green: m["green"],
			},
		}

		if result.isPossible() {
			score += result.index + 1
		}
	}

	return score
}

type Game struct {
	index int
	turn  Turn
}

func (game *Game) isPossible() bool {
	if game.turn.red <= 12 && game.turn.blue <= 13 && game.turn.green <= 14 {
		return true
	}
	return false
}

type Turn struct {
	blue  int
	red   int
	green int
}
