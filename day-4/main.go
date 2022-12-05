package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day-4/input.txt")
	scanner := bufio.NewScanner(file)
  complete := 0
  partial := 0

	defer file.Close()

	for scanner.Scan() {
		elfs := strings.Split(scanner.Text(), ",")
    first := strings.Split(elfs[0], "-")
    second := strings.Split(elfs[1], "-")
		firstStart, _ := strconv.Atoi(first[0])
		firstEnd, _ := strconv.Atoi(first[1])
		secondStart, _ := strconv.Atoi(second[0])
		secondEnd, _ := strconv.Atoi(second[1])

		if (firstStart >= secondStart && firstStart <= secondEnd) &&
			(firstEnd >= secondStart && firstEnd <= secondEnd) ||
			(secondStart >= firstStart && secondStart <= firstEnd) &&
			(secondEnd >= firstStart && secondEnd <= firstEnd) {
      complete++
		}

		if (firstStart >= secondStart && firstStart <= secondEnd) ||
			(firstEnd >= secondStart && firstEnd <= secondEnd) ||
			(secondStart >= firstStart && secondStart <= firstEnd) ||
			(secondEnd >= firstStart && secondEnd <= firstEnd) {
			partial++
		}
	}

	fmt.Printf("Part one: %d\n", complete)
	fmt.Printf("Part two: %d\n", partial)
}
