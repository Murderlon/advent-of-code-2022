package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day-8/input.txt")
	scanner := bufio.NewScanner(file)
	trees := make([][]int, 0)
	visible_trees := 0
  max_view := 0

	defer file.Close()

	for scanner.Scan() {
		line := make([]int, 0)
		for _, tree := range scanner.Text() {
			height, _ := strconv.Atoi(string(tree))
			line = append(line, height)
		}
		trees = append(trees, line)
	}

	for y, line := range trees {
		if y-1 < 0 || y+1 > len(trees)-1 {
			visible_trees += len(line)
			continue
		}
		for x := range line {
			if x == 0 || x == len(line)-1 {
				visible_trees++
				continue
			}
			if visible(trees, x, y) {
				visible_trees++
			}
      if n := count(trees, x, y); n > max_view {
        max_view = n
      }
		}
	}

	fmt.Println("Part one:", visible_trees)
	fmt.Println("Part two:", max_view)
}

func count(grid [][]int, x, y int) int {
	left := 0
	right := 0
	top := 0
	bottom := 0
	height := grid[y][x]

	for i := x - 1; i >= 0; i-- {
		left++
		if grid[y][i] >= height {
			break
		}
	}
	for i := x + 1; i < len(grid[y]); i++ {
		right++
		if grid[y][i] >= height {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		top++
		if grid[i][x] >= height {
			break
		}
	}
	for i := y + 1; i < len(grid[y]); i++ {
		bottom++
		if grid[i][x] >= height {
			break
		}
	}

	return left * right * top * bottom
}

func visible(grid [][]int, x, y int) bool {
	left := true
	right := true
	top := true
	bottom := true
	height := grid[y][x]

	for i := range grid[y] {
		if i < x && grid[y][i] >= height {
			left = false
		}
		if i > x && grid[y][i] >= height {
			right = false
		}
	}
	for i := range grid {
		if i < y && grid[i][x] >= height {
			top = false
		}
		if i > y && grid[i][x] >= height {
			bottom = false
		}
	}

	return left || right || top || bottom
}
