package main

// Solution for AoC 2022 Day 6
// https://adventofcode.com/2022/day/6

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadByteSteam(r io.Reader) []rune {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	return []rune(scanner.Text())
}

func DistanceToFirstMarker(datastream []rune, markerLength int, isValid func([]rune) bool) int {
	for i := 0; i < len(datastream)-markerLength; i++ {
		marker := datastream[i : i+markerLength]
		if isValid(marker) {
			return i + markerLength
		}
	}
	return -1
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	datastream := ReadByteSteam(handle)

	markerValidationFunction := func(marker []rune) bool {
		set := map[rune]struct{}{}
		for _, r := range marker {
			set[r] = struct{}{}
		}
		return len(set) == len(marker)
	}

	fmt.Printf("Part 1: %d\n", DistanceToFirstMarker(datastream, 4, markerValidationFunction))
	fmt.Printf("Part 2: %d\n", DistanceToFirstMarker(datastream, 14, markerValidationFunction))
}
