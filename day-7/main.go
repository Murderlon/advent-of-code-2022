package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	disk_space   = 70000000
	space_needed = 30000000
)

type Node struct {
	name     string
	parent   *Node
	size     int
	children []*Node
}

func (node *Node) visit(visitor func(*Node)) {
	visitor(node)
	for _, child := range node.children {
		child.visit(visitor)
	}
}

func (node *Node) sum() int {
	for _, child := range node.children {
		node.size += child.sum()
	}
	return node.size
}

func main() {
	file, _ := os.Open("day-7/input.txt")
	scanner := bufio.NewScanner(file)
	root := Node{name: "/", children: []*Node{}}
	cwd := &root
	sum_of_dirs, rm_dir_size := 0, disk_space

	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		instructions := strings.Fields(line)

		switch {
		case instructions[0] == "$":
			if instructions[1] == "cd" {
				arg := instructions[2]
				switch arg {
				case "/":
					cwd = &root
				case "..":
					cwd = cwd.parent
				default:
					for _, child := range cwd.children {
						if child.name == arg {
							cwd = child
						}
					}
				}
			}
		case instructions[0] == "dir":
			cwd.children = append(cwd.children, &Node{name: instructions[1], parent: cwd})
		default:
			size, _ := strconv.Atoi(instructions[0])
			cwd.children = append(cwd.children, &Node{name: instructions[1], size: size, parent: cwd})
		}
	}

	root.sum()
	remainder := space_needed - (disk_space - root.size)
	root.visit(func(node *Node) {
		if len(node.children) > 0 && node.size <= 100000 {
			sum_of_dirs += node.size
		}
		if len(node.children) > 0 && node.size >= remainder && node.size < rm_dir_size {
			rm_dir_size = node.size
		}
	})

	fmt.Println("Part one:", sum_of_dirs)
	fmt.Println("Part two:", rm_dir_size)
}
