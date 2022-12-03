package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	input, err := os.ReadFile("day_03/input.txt")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	rucksacks := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println(sumPriority(rucksacks))
	fmt.Println(sumGroups(rucksacks))
}

func sumPriority(rucksacks []string) int {
	sum := 0

	for _, rs := range rucksacks {
		l := len(rs) / 2
		cp1, cp2 := rs[0:l], rs[l:]

		for _, c := range cp1 {
			if slices.Contains([]rune(cp2), c) {
				sum += priority(c)
				break
			}
		}
	}

	return sum
}

func sumGroups(rucksacks []string) int {
	sum := 0

	for i := 0; i < len(rucksacks)/3; i++ {
		group := rucksacks[i*3 : i*3+3]

		for _, c := range group[0] {
			if slices.Contains([]rune(group[1]), c) && slices.Contains([]rune(group[2]), c) {
				sum += priority(c)
				break
			}
		}
	}

	return sum
}

func priority(char rune) int {
	if char < 'a' {
		return int(char - 38)
	} else {
		return int(char - 96)
	}
}
