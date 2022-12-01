package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("Part one: %d\n", one())
	fmt.Printf("Part two: %d\n", two())
}

func one() int {
	file, _ := os.Open("day-1/input.txt")
	scanner := bufio.NewScanner(file)
	max := 0
	current := 0

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if current > max {
				max = current
			}
			current = 0
		} else {
			number, _ := strconv.Atoi(line)
			current += number
		}
	}

	return max
}

func two() int {
	file, _ := os.Open("day-1/input.txt")
	scanner := bufio.NewScanner(file)
	slice := make([]int, 0)
	current := 0

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			slice = append(slice, current)
			current = 0
		} else {
			number, _ := strconv.Atoi(line)
			current += number
		}
	}
	slice = append(slice, current)
	sort.Ints(slice)
	top := slice[len(slice)-3:]
	return top[0] + top[1] + top[2]
}
