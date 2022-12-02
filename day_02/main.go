package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	plMap = map[string]int{
		// Rock
		"X": 1,
		// Paper
		"Y": 2,
		// Scissors
		"Z": 3,
	}
	enMap = map[string]int{
		// Rock
		"A": 1,
		// Paper
		"B": 2,
		// Scissors
		"C": 3,
	}
	scoreMap = map[string]map[string]int{
		// Lose
		"X": {"A": 3, "B": 1, "C": 2},
		// Draw
		"Y": {"A": 1, "B": 2, "C": 3},
		// Win
		"Z": {"A": 2, "B": 3, "C": 1},
	}
	resMap = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
)

func calcScore1(matches []string) int {
	score := 0

	for _, match := range matches {

		p := strings.Split(match, " ")
		pl, en := plMap[p[1]], enMap[p[0]]

		score += pl

		if pl == 1 && en == 3 || pl == 2 && en == 1 || pl == 3 && en == 2 {
			// Win
			score += 6
		} else if pl == en {
			// Draw
			score += 3
		}

	}

	return score
}

func calcScore2(matches []string) int {
	score := 0

	for _, match := range matches {
		p := strings.Split(match, " ")
		en, res := p[0], p[1]
		score += scoreMap[res][en] + resMap[res]
	}

	return score
}

func main() {
	input, err := os.ReadFile("day_02/input.txt")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	matches := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println(calcScore1(matches))
	fmt.Println(calcScore2(matches))
}
