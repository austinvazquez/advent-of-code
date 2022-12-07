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

type FileEntry struct {
	name        string
	size        int
	isDirectory bool
	parent      *FileEntry
	children    map[string]*FileEntry
}

type FileEntryOpt func(*FileEntry)

func WithSize(size int) FileEntryOpt {
	return func(entry *FileEntry) {
		entry.size = size
	}
}

func WithChild(child *FileEntry) FileEntryOpt {
	return func(entry *FileEntry) {
		if entry.isDirectory {
			entry.children[child.name] = entry
		}
	}
}

func NewFileEntry(name string, isDirectory bool, parent *FileEntry, opts ...FileEntryOpt) *FileEntry {
	entry := &FileEntry{
		name:        name,
		isDirectory: isDirectory,
		parent:      parent,
		children:    make(map[string]*FileEntry),
	}

	for _, opt := range opts {
		opt(entry)
	}

	return entry
}

func (f *FileEntry) AddFile(name string, size int) {
	f.children[name] = &FileEntry{name: name, size: size, parent: f}
}

func (f *FileEntry) AddDirectory(name string) {
	f.children[name] = &FileEntry{
		name:        name,
		isDirectory: true,
		parent:      f,
		children:    make(map[string]*FileEntry),
	}
}

func (f *FileEntry) Size() int {
	if f.isDirectory {
		var size int
		for _, v := range f.children {
			size += v.Size()
		}
		return size
	}
	return f.size
}

func ReadFilesystem(r io.Reader) *FileEntry {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var root *FileEntry
	var currentDirectory *FileEntry

	for i := 0; i < len(lines); {
		line := lines[i]

		if strings.Contains(line, "cd ..") {
			if currentDirectory.parent != nil {
				currentDirectory = currentDirectory.parent
			}
			i++
		} else if strings.Contains(line, "cd") {
			fields := strings.Fields(line)
			name := fields[2]

			var entry *FileEntry
			if root == nil {
				root = NewFileEntry(name, true, currentDirectory)
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
					entry.AddDirectory(name)
				} else {
					size, err := strconv.Atoi(fields[0])
					check(err)

					entry.AddFile(name, size)
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

func Traverse(root *FileEntry) []*FileEntry {
	nodes := []*FileEntry{root}

	for _, v := range root.children {
		if v.isDirectory {
			nodes = append(nodes, Traverse(v)...)
		} else {
			nodes = append(nodes, v)
		}
	}
	return nodes
}

func Filter(array []*FileEntry, f func(*FileEntry) bool) []*FileEntry {
	var filtered []*FileEntry
	for _, n := range array {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func Reduce(array []*FileEntry, f func(*FileEntry) int) int {
	accumulator := 0

	for _, n := range array {
		accumulator += f(n)
	}

	return accumulator
}

func Min(array []*FileEntry, f func(*FileEntry) int) int {
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

	atMost := func(size int) func(*FileEntry) bool {
		return func(entry *FileEntry) bool {
			return entry.isDirectory && entry.Size() < size
		}
	}

	atLeast := func(size int) func(*FileEntry) bool {
		return func(entry *FileEntry) bool {
			return entry.isDirectory && entry.Size() >= size
		}
	}

	size := func(entry *FileEntry) int {
		return entry.Size()
	}

	usedSpace := root.Size()
	freeSpace := 70000000 - usedSpace
	fmt.Printf("Part 1: %d\n", Reduce(Filter(Traverse(root), atMost(100000)), size))
	fmt.Printf("Part 2: %d\n", Min(Filter(Traverse(root), atLeast(30000000-freeSpace)), size))
}
