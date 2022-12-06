package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

var createRegex = regexp.MustCompile(`\[[A-Z]]`)
var instructionRegex = regexp.MustCompile(`move`)

func main() {
	file, _ := os.Open("day-5/input.txt")
	scanner := bufio.NewScanner(file)
	crateLines := make([]string, 0)
	one := make([][]string, 0)
	two := make([][]string, 0)

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case createRegex.MatchString(line):
			crateLines = append(crateLines, line)
		case len(line) == 0:
			stack := make([]string, 0)
			for x := range crateLines[0] {
				for y := len(crateLines) - 1; y >= 0; y-- {
					if unicode.IsLetter(rune(crateLines[y][x])) {
						stack = append(stack, string(crateLines[y][x]))
					}
				}
				if len(stack) > 0 {
					one = append(one, stack)
					stack = make([]string, 0)
				}
			}
			deepClone(&two, &one)
		case instructionRegex.MatchString(line):
			s := strings.Fields(line)
			amount, _ := strconv.Atoi(s[1])
			from, _ := strconv.Atoi(s[3])
			to, _ := strconv.Atoi(s[5])
			from--
			to--

			one[to] = append(one[to], reverse(one[from][len(one[from])-amount:])...)
			one[from] = slices.Delete(one[from], len(one[from])-amount, len(one[from]))
			two[to] = append(two[to], two[from][len(two[from])-amount:]...)
			two[from] = slices.Delete(two[from], len(two[from])-amount, len(two[from]))
		default:
			continue
		}
	}

	fmt.Println("Part one:", message(one))
	fmt.Println("Part two:", message(two))
}

func deepClone(dst *[][]string, source *[][]string) {
  d := *dst
  s := *source
	for i := range s {
		d = append(d, make([]string, len(s[i])))
		copy(d[i], s[i])
	}
  *dst = d
}

func message(s [][]string) string {
	message := ""
	for i := range s {
		message += s[i][len(s[i])-1]
	}
	return message
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
