package main

// Solution for AoC 2022 Day 7
// https://adventofcode.com/2022/day/7

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	name     string
	size     int
	parent   *Node
	children map[string]*Node
}

func NewNode(name string, parent *Node) *Node {
	self := &Node{
		name:     name,
		children: make(map[string]*Node),
	}
	if parent == nil {
		parent = self
	}
	self.parent = parent
	return self
}

func (n *Node) AddLeaf(name string, size int) {
	n.children[name] = &Node{
		name:   name,
		size:   size,
		parent: n,
	}
}

func (n *Node) AddBranch(name string) {
	n.children[name] = &Node{
		name:     name,
		parent:   n,
		children: make(map[string]*Node),
	}
}

func (n *Node) Size() int {
	if n.IsBranch() {
		var size int
		for _, v := range n.children {
			size += v.Size()
		}
		return size
	}
	return n.size
}

func (n *Node) IsBranch() bool {
	return n.children != nil && len(n.children) > 0
}

func ReadFilesystem(r io.Reader) *Node {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var root *Node
	var currentDirectory *Node

	for i := 0; i < len(lines); {
		line := lines[i]

		if strings.Contains(line, "cd ..") {
			currentDirectory = currentDirectory.parent
			i++
		} else if strings.Contains(line, "cd") {
			fields := strings.Fields(line)
			name := fields[2]

			var entry *Node
			if root == nil {
				root = NewNode(name, currentDirectory)
				entry = root
			} else {
				entry = currentDirectory.children[name]
			}
			currentDirectory = entry

			i += 2
			line = lines[i]

			for !strings.Contains(line, "$") {
				fields := strings.Fields(line)
				name := fields[1]

				if strings.Contains(line, "dir") {
					entry.AddBranch(name)
				} else {
					size, err := strconv.Atoi(fields[0])
					check(err)

					entry.AddLeaf(name, size)
				}

				i += 1

				if i >= len(lines) {
					break
				}
				line = lines[i]
			}
		} else {
			i++
		}
	}

	return root
}

func Traverse(root *Node) []*Node {
	nodes := []*Node{root}

	for _, v := range root.children {
		if v.IsBranch() {
			nodes = append(nodes, Traverse(v)...)
		} else {
			nodes = append(nodes, v)
		}
	}

	return nodes
}

func Filter(array []*Node, f func(*Node) bool) []*Node {
	var filtered []*Node

	for _, n := range array {
		if f(n) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}

func Reduce(array []*Node, f func(*Node) int) int {
	accumulator := 0

	for _, n := range array {
		accumulator += f(n)
	}

	return accumulator
}

func Min(array []*Node, f func(*Node) int) int {
	min := f(array[0])

	for i := 1; i < len(array); i++ {
		if fn := f(array[i]); fn < min {
			min = fn
		}
	}

	return min
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	root := ReadFilesystem(handle)

	atMost := func(size int) func(*Node) bool {
		return func(entry *Node) bool {
			return entry.IsBranch() && entry.Size() < size
		}
	}

	atLeast := func(size int) func(*Node) bool {
		return func(entry *Node) bool {
			return entry.IsBranch() && entry.Size() >= size
		}
	}

	size := func(entry *Node) int {
		return entry.Size()
	}

	fmt.Printf("Part 1: %d\n", Reduce(Filter(Traverse(root), atMost(100000)), size))

	usedSpace := root.Size()
	freeSpace := 70000000 - usedSpace
	fmt.Printf("Part 2: %d\n", Min(Filter(Traverse(root), atLeast(30000000-freeSpace)), size))
}
