// https://adventofcode.com/2022/day/1

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	buf, err := os.ReadFile("day_01/input.txt")
	if err != nil {
		fmt.Println("An error occured")
		return
	}

	elves, err := parse(string(buf))
	if err != nil {
		fmt.Println("An error occured")
		return
	}

	fmt.Println(getMax(elves))

	fmt.Println(getMax3(elves))
}

func parse(str string) ([]int, error) {
	lines := strings.Split(strings.TrimSpace(str), "\n")

	elves := []int{0}
	i := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, 0)
			i++
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		elves[i] += num
	}

	return elves, nil
}

func getMax(elves []int) int {
	max := elves[0]
	for _, elf := range elves {
		if elf > max {
			max = elf
		}
	}

	return max
}

func getMax3(elves []int) int {
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})
	sum := 0
	for _, elf := range elves[:3] {
		sum += elf
	}
	return sum
}
