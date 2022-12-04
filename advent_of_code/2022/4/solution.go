package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Solution for AoC 2022 Day 4
// https://adventofcode.com/2022/day/4

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Range struct {
	start int
	end   int
}

func NewRange(rangeString string) *Range {
	ints := strings.Split(rangeString, "-")

	start, err := strconv.Atoi(ints[0])
	check(err)

	end, err := strconv.Atoi(ints[1])
	check(err)

	return &Range{start, end}
}

func (r *Range) Contains(other *Range) bool {
	return other.start >= r.start && other.end <= r.end
}

func (r *Range) Overlaps(other *Range) bool {
	return r.start <= other.end && r.end >= other.start
}

type Assignment struct {
	first  *Range
	second *Range
}

func NewAssignment(written string) *Assignment {
	assignment := strings.Split(written, ",")
	return &Assignment{NewRange(assignment[0]), NewRange(assignment[1])}
}

func (a *Assignment) Subsets() bool {
	return a.first.Contains(a.second) || a.second.Contains(a.first)
}

func (a *Assignment) Overlaps() bool {
	return a.first.Overlaps(a.second)
}

func ReadAssignments(r io.ReadCloser) []*Assignment {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var assignments []*Assignment
	for scanner.Scan() {
		assignments = append(assignments, NewAssignment(scanner.Text()))
	}

	return assignments
}

func Reduce(assignments []*Assignment, f func(*Assignment) int) int {
	accumulator := 0

	for _, assignment := range assignments {
		accumulator += f(assignment)
	}

	return accumulator
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	assignments := ReadAssignments(handle)

	countSubsets := func(assignment *Assignment) int {
		if assignment.Subsets() {
			return 1
		}
		return 0
	}

	countOverlaps := func(assignment *Assignment) int {
		if assignment.Overlaps() {
			return 1
		}
		return 0
	}

	fmt.Printf("Part 1: %d\n", Reduce(assignments, countSubsets))
	fmt.Printf("Part 2: %d\n", Reduce(assignments, countOverlaps))
}
