package main

// Solution for AoC 2021 Day 1
// https://adventofcode.com/2021/day/1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var ints []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return ints, err
		}
		ints = append(ints, i)
	}

	return ints, scanner.Err()
}

func CountIncreasesWithSlidingWindow(ints []int, windowSize int) int {
	/*
	   Complexity analysis:
	   Time: O(n)
	   Space: O(1)
	*/
	count := 0

	for i := 0; i < len(ints)-windowSize; i++ {
		if ints[i+windowSize] > ints[i] {
			count++
		}
	}

	return count
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)

	defer handle.Close()

	nums, err := ReadInts(handle)
	check(err)

	fmt.Printf("Part 1: %d\n", CountIncreasesWithSlidingWindow(nums, 1))
	fmt.Printf("Part 2: %d\n", CountIncreasesWithSlidingWindow(nums, 3))
}
