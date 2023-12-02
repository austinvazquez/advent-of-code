package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Solution for AoC 2023 Day 0
// https://adventofcode.com/2023/day/0

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	lines, err := ReadLines(file)
	check(err)

	identity := func(i int) int {
		return i
	}

	fmt.Printf("Step 1: %d", Reduce(Map(lines, Transform), identity))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func Reduce(nums []int, f func(int) int) int {
	accumulator := 0

	for _, n := range nums {
		accumulator += f(n)
	}

	return accumulator
}

func Map(lines []string, f func(string) int) []int {
	nums := []int{}

	for _, line := range lines {
		nums = append(nums, f(line))
	}

	return nums
}

func Transform(line string) int {
	firstDigit, secondDigit := 0, 0

	for _, r := range line {
		if unicode.IsDigit(r) {
			firstDigit = int(r - '0')
			break
		}
	}

	for i := range line {
		r := []rune(line)[len(line)-1-i]
		if unicode.IsDigit(r) {
			secondDigit = int(r - '0')
			break
		}
	}

	return firstDigit<<1 + secondDigit
}
