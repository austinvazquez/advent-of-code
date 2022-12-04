package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// Solution for AoC 2022 Day 3
// https://adventofcode.com/2022/day/3

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadRucksacks(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rucksacks []string
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	return rucksacks
}

func Split(rucksack string) (string, string) {
	var left, right []rune

	for i, ch := range rucksack {
		if i < len(rucksack)/2 {
			left = append(left, ch)
		} else {
			right = append(right, ch)
		}
	}

	return string(left), string(right)
}

func Group(rusksacks []string) [][]string {
	var groups [][]string

	for i := 0; i < len(rusksacks); i += 3 {
		groups = append(groups, []string{rusksacks[i], rusksacks[i+1], rusksacks[i+2]})
	}

	return groups
}

func FindIntersectionOfCompartments(lhs, rhs string) rune {
	for _, ch := range lhs {
		if strings.Contains(rhs, string(ch)) {
			if unicode.IsUpper(ch) {
				return ch - 38
			}

			return ch - 96
		}
	}
	return -1
}

func FindIntersectionOfGroup(group []string) rune {
	for _, ch := range group[0] {
		if strings.Contains(group[1], string(ch)) && strings.Contains(group[2], string(ch)) {
			if unicode.IsUpper(ch) {
				return ch - 38
			}

			return ch - 96
		}
	}
	return -1
}

func GetPriorityFromRusksack(rucksack string) int {
	return int(FindIntersectionOfCompartments(Split(rucksack)))
}

func GetPriorityFromGroup(group []string) int {
	return int(FindIntersectionOfGroup(group))
}

func Reduce(array []string, f func(string) int) int {
	accumulator := 0

	for _, n := range array {
		accumulator += f(n)
	}

	return accumulator
}

func ReduceGroup(array []string, f func([]string) int) int {
	accumulator := 0

	for _, n := range Group(array) {
		accumulator += f(n)
	}

	return accumulator
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	rucksacks := ReadRucksacks(handle)

	fmt.Printf("Part 1: %d\n", Reduce(rucksacks, GetPriorityFromRusksack))
	fmt.Printf("Part 2: %d\n", ReduceGroup(rucksacks, GetPriorityFromGroup))
}
