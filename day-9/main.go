package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
  x, y int
}

type Rope struct {
  head, tail Point
}

// x x x x x
// x T T T x
// x T H T x
// x T T T x
// x x x x x
func (r *Rope) isCollinear() bool {
	dx := float64(r.tail.x - r.head.x)
	dy := float64(r.tail.y - r.head.y)

	if math.Abs(dy) <= 1 && math.Abs(dx) <= 1 {
		return true
	}
	return false
}

func (r *Rope) Move(direction string) {
	switch direction {
	case "R":
		r.head.x++
		if r.isCollinear() {
			return
		}
		r.tail.x++
		r.tail.y = r.head.y
	case "L":
		r.head.x--
		if r.isCollinear() {
			return
		}
		r.tail.x--
		r.tail.y = r.head.y
	case "U":
		r.head.y++
		if r.isCollinear() {
			return
		}
		r.tail.y++
		r.tail.x = r.head.x
	case "D":
		r.head.y--
		if r.isCollinear() {
			return
		}
		r.tail.y--
		r.tail.x = r.head.x
	}
}

func main() {
	file, _ := os.Open("day-9/example.txt")
	scanner := bufio.NewScanner(file)
	one := make(map[Point]int)
	two := make(map[Point]int)
	rope := Rope{}
	ropes := make([]Rope, 10)

	defer file.Close()

	for scanner.Scan() {
		instructions := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(instructions[1])
		direction := instructions[0]

		for i := 0; i < n; i++ {
			rope.Move(direction)
			one[rope.tail] += 1

      // Part two
      for j := range ropes {
        if i-j >= 0 {
          ropes[j].Move(direction)
        }
        if j < len(ropes)-1 {
          ropes[j+1].head.x = ropes[j].tail.x
          ropes[j+1].head.y = ropes[j].tail.y
        }
        two[ropes[9].tail] += 1
      }
		}
	}

	fmt.Println("Part one:", len(one))
	fmt.Println("Part two:", len(two))
}
