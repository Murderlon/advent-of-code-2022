package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rock    = 1
	paper   = 2
	scissor = 3

	win  = 6
	draw = 3
)

var alias = map[string]string{
	"X": "rock",
	"A": "rock",
	"Y": "paper",
	"B": "paper",
	"Z": "scissor",
	"C": "scissor",
}

func main() {
	fmt.Printf("Part one: %d\n", one())
	fmt.Printf("Part two: %d\n", two())
}

func one() int {
	file, _ := os.Open("day-2/input.txt")
	scanner := bufio.NewScanner(file)
	game := map[string]map[string]int{
		"rock":    {"rock": rock + draw, "paper": paper + win, "scissor": scissor},
		"paper":   {"rock": rock, "paper": paper + draw, "scissor": scissor + win},
		"scissor": {"rock": rock + win, "paper": paper, "scissor": scissor + draw},
	}
	total := 0

	defer file.Close()

	for scanner.Scan() {
		picks := strings.Fields(scanner.Text())
		them := alias[picks[0]]
		me := alias[picks[1]]

		total += game[them][me]
	}

	return total
}

func two() int {
	file, _ := os.Open("day-2/input.txt")
	scanner := bufio.NewScanner(file)
	mustWin := map[string]int{"rock": paper + win, "paper": scissor + win, "scissor": rock + win}
	mustDraw := map[string]int{"rock": rock + draw, "paper": paper + draw, "scissor": scissor + draw}
	mustLoose := map[string]int{"rock": scissor, "paper": rock, "scissor": paper}
  game := map[string]map[string]int{"Z": mustWin, "Y": mustDraw, "X": mustLoose}
	total := 0

	defer file.Close()

	for scanner.Scan() {
		picks := strings.Fields(scanner.Text())
		them := alias[picks[0]]
		me := picks[1]

    total += game[me][them]
	}

	return total
}
