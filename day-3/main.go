package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("Part one: %d\n", one())
	fmt.Printf("Part two: %d\n", two())
}

func priority(char rune) int {
	if unicode.IsUpper(char) {
		return int(char) - 38
	} else {
		return int(char) - 96
	}
}

func one() int {
	file, _ := os.Open("day-3/input.txt")
	scanner := bufio.NewScanner(file)
	total := 0

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		half := len(line) / 2
		dupes := make(map[rune]bool)

		for index, char := range line {
			i := strings.IndexRune(line, char)

			if _, ok := dupes[char]; !ok && index >= half && i != -1 && i < half {
				total += priority(char)
				dupes[char] = true
			}
		}
	}

	return total
}

func two() int {
	file, _ := os.Open("day-3/input.txt")
	scanner := bufio.NewScanner(file)
	group := ""
	index := 0
	total := 0

	defer file.Close()

	for scanner.Scan() {
		uniques := ""
		for _, char := range scanner.Text() {
			if strings.ContainsRune(uniques, char) {
				continue
			}
			uniques += string(char)
		}
		group += uniques

		index++
		if index > 2 {
			for _, char := range group {
				if strings.Count(group, string(char)) == 3 {
					total += priority(char)
					break
				}
			}

			index = 0
			group = ""
		}
	}

	return total
}
