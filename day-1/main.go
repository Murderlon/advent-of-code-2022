package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
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
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	top := slice[:3]

	fmt.Printf("Part one: %d\n", top[0])
	fmt.Printf("Part two: %d\n", top[0]+top[1]+top[2])
}
