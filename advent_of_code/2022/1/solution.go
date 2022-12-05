package main

// Solution for AoC 2022 Day 1
// https://adventofcode.com/2022/day/1

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
	scanner.Split(bufio.ScanLines)

	var ints []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			i = 0
		}

		ints = append(ints, i)
	}

	return ints, scanner.Err()
}

func Reduce(nums []int, f func(int) int) int {
	accumulator := 0

	for _, n := range nums {
		accumulator += f(n)
	}

	return accumulator
}

func FindTopNCalorieCounts(nums []int, top int) []int {
	/*
	   Complexity analysis:
	   Time: O(n)
	   Space: O(1)
	*/
	calories := make([]int, top)

	addCalorieCount := func(calorie int) {
		for i, c := range calories {
			if calorie > c {
				calories[i] = calorie
				calorie = c
			}
		}
	}

	calorieCount := 0
	for _, num := range nums {
		if num == 0 {
			addCalorieCount(calorieCount)
			calorieCount = 0
		} else {
			calorieCount += num
		}
	}

	return calories
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	nums, err := ReadInts(handle)
	check(err)

	reflect := func(n int) int {
		return n
	}

	fmt.Printf("Part 1: %d\n", Reduce(FindTopNCalorieCounts(nums, 1), reflect))
	fmt.Printf("Part 2: %d\n", Reduce(FindTopNCalorieCounts(nums, 3), reflect))
}
