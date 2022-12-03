package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	A = 1
	B = 2
	C = 3

	win  = 6
	draw = 3
)

var alias = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

var game = map[string]map[string]int{
	"A": {"A": A + draw, "B": B + win, "C": C},
	"B": {"A": A, "B": B + draw, "C": C + win},
	"C": {"A": A + win, "B": B, "C": C + draw},
}
var mustWin = map[string]string{"A": "B", "B": "C", "C": "A"}
var mustDraw = map[string]string{"A": "A", "B": "B", "C": "C"}
var mustLoose = map[string]string{"A": "C", "B": "A", "C": "B"}

func main() {
	fmt.Printf("Part one: %d\n", one())
	fmt.Printf("Part two: %d\n", two())
}

func one() int {
	file, _ := os.Open("day-2/input.txt")
	scanner := bufio.NewScanner(file)
	total := 0

	defer file.Close()

	for scanner.Scan() {
		picks := strings.Fields(scanner.Text())
		them := picks[0]
		me := alias[picks[1]]

		total += game[them][me]
	}

	return total
}

func two() int {
	file, _ := os.Open("day-2/input.txt")
	scanner := bufio.NewScanner(file)
	instruction := map[string]map[string]string{"Z": mustWin, "Y": mustDraw, "X": mustLoose}
	total := 0

	defer file.Close()

	for scanner.Scan() {
		picks := strings.Fields(scanner.Text())
		them := picks[0]
		me := instruction[picks[1]][them]

		total += game[them][me]
	}

	return total
}
