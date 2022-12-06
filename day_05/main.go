package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ship [9][]string

func main() {

	input, err := os.ReadFile("day_05/input.txt")
	if err != nil {
		fmt.Println("Something went wrong")
		return
	}

	parts := strings.Split(string(input), "\n\n")

	ship := parseShip(parts[0])

	ship1 := move1(ship, parts[1])
	printShip(ship1)

	ship2 := move2(ship, parts[1])
	printShip(ship2)

}

func parseShip(shipString string) ship {

	ship := ship{}

	shipLines := strings.Split(shipString, "\n")
	for i, j := 0, len(shipLines)-1; i < j; i, j = i+1, j-1 {
		shipLines[i], shipLines[j] = shipLines[j], shipLines[i]
	}

	for _, l := range shipLines[1:] {
		for i := 0; i < len(l); i += 4 {
			stackI := i / 4
			cargo := string(l[i+1])
			if cargo != " " {
				ship[stackI] = append(ship[stackI], cargo)
			}
		}
	}

	return ship
}

func printShip(ship ship) {
	out := ""
	for _, s := range ship {
		out += s[len(s)-1]
	}

	fmt.Println(out)
}

func move1(ship ship, instructions string) ship {

	instLines := strings.Split(strings.TrimSpace(instructions), "\n")

	for _, inst := range instLines {

		parts := strings.Split(strings.TrimSpace(inst), " ")

		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		from--
		to, _ := strconv.Atoi(parts[5])
		to--

		for i := 0; i < count; i++ {

			l := len(ship[from])

			cargo := ship[from][l-1]
			ship[from] = ship[from][:l-1]
			ship[to] = append(ship[to], cargo)

		}

	}

	return ship
}

func move2(ship ship, instructions string) ship {

	instLines := strings.Split(strings.TrimSpace(instructions), "\n")

	for _, inst := range instLines {

		parts := strings.Split(strings.TrimSpace(inst), " ")

		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		from--
		to, _ := strconv.Atoi(parts[5])
		to--

		l := len(ship[from])

		cargo := ship[from][l-count : l]

		ship[from] = ship[from][:l-count]
		ship[to] = append(ship[to], cargo...)

	}

	return ship
}
