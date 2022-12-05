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
	sections := make([][]int, 0)
  complete := 0
  partial := 0

	defer file.Close()

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")

		for _, pair := range pairs {
			nums := strings.Split(pair, "-")
			start, _ := strconv.Atoi(nums[0])
			end, _ := strconv.Atoi(nums[1])
			sections = append(sections, []int{start, end})
		}
	}

	for i := 1; i < len(sections); i += 2 {
		firstStart := sections[i-1][0]
		firstEnd := sections[i-1][1]
		secondStart := sections[i][0]
		secondEnd := sections[i][1]

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
