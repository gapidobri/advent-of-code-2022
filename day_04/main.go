package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("day_04/input.txt")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	pairs := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println(getContainCount(pairs))
	fmt.Println(getOverlapCount(pairs))

}

func getContainCount(pairs []string) int {

	count := 0

	for _, pair := range pairs {

		asig := strings.Split(pair, ",")
		a1p := strings.Split(asig[0], "-")
		a2p := strings.Split(asig[1], "-")

		a1n1, _ := strconv.Atoi(a1p[0])
		a1n2, _ := strconv.Atoi(a1p[1])
		a2n1, _ := strconv.Atoi(a2p[0])
		a2n2, _ := strconv.Atoi(a2p[1])

		if a1n1 <= a2n1 && a1n2 >= a2n2 || a2n1 <= a1n1 && a2n2 >= a1n2 {
			count++
		}

	}

	return count
}

func getOverlapCount(pairs []string) int {
	count := 0

	for _, pair := range pairs {

		asig := strings.Split(pair, ",")
		a1p := strings.Split(asig[0], "-")
		a2p := strings.Split(asig[1], "-")

		a1n1, _ := strconv.Atoi(a1p[0])
		a1n2, _ := strconv.Atoi(a1p[1])
		a2n1, _ := strconv.Atoi(a2p[0])
		a2n2, _ := strconv.Atoi(a2p[1])

		if a1n1 <= a2n1 && a1n2 >= a2n1 ||
			a2n1 <= a1n1 && a2n2 >= a1n1 ||
			a1n1 <= a2n2 && a1n2 >= a2n2 ||
			a2n1 <= a1n2 && a2n2 >= a1n2 {
			count++
		}

	}

	return count
}
