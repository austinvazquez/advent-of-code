package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Solution for AoC 2023 Day 1
// https://adventofcode.com/2023/day/1

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	lines, err := ReadLines(file)
	check(err)

	sum := func(l, r int) int {
		return l + r
	}

	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Printf("Step 1: %d\n", Reduce(
		Map(lines, func(t string) int {
			return Callibrate(t, digits)
		}), sum, 0))

	digits = append(digits, []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}...)
	fmt.Printf("Step 2: %d\n", Reduce(
		Map(lines, func(t string) int {
			return Callibrate(t, digits)
		}), sum, 0))
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

func Reduce[T, M any](collection []T, f func(M, T) M, initial M) M {
	accumulator := initial

	for _, c := range collection {
		accumulator = f(accumulator, c)
	}

	return accumulator
}

func Map[T, M any](collection []T, f func(T) M) []M {
	mc := []M{}

	for _, c := range collection {
		mc = append(mc, f(c))
	}

	return mc
}

var textToInt = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Callibrate(text string, digits []string) int {
	mappedDigits := map[string]int{}

	for _, digit := range digits {
		if strings.Contains(text, digit) {
			mappedDigits[digit] = strings.Index(text, digit)
		}
	}

	firstDigit, index := 0, len(text)

	for k, v := range mappedDigits {
		if v < index {
			index = v
			firstDigit = textToInt[k]
		}
	}

	mappedDigits = map[string]int{}

	for _, digit := range digits {
		if strings.Contains(text, digit) {
			mappedDigits[digit] = strings.LastIndex(text, digit)
		}
	}

	lastDigit, index := 0, -1

	for k, v := range mappedDigits {
		if v > index {
			index = v
			lastDigit = textToInt[k]
		}
	}

	return (firstDigit * 10) + lastDigit
}
