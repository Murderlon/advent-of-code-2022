package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const sequence = 14

func main() {
	file, _ := os.Open("day-6/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		for i := range line {
			if i >= sequence-1 {
				chars := line[i-sequence+1 : i+1]
				found := true
				for j, char := range chars {
					if r := strings.IndexRune(chars, char); r != j {
						found = false
					}
				}
				if found {
					fmt.Println(i + 1)
					break
				}
			}
		}
	}
}
