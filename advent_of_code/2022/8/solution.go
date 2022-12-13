package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Solution for AoC 2022 Day 8
// https://adventofcode.com/2022/day/8

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadForest(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var forest [][]int
	for scanner.Scan() {
		var line []int
		for _, tree := range []rune(scanner.Text()) {
			line = append(line, int(tree)-48)
		}
		forest = append(forest, line)
	}

	return forest
}

func Count(array [][]int, counts func([][]int, int, int) bool) int {
	count := 0

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			if counts(array, i, j) {
				count++
			}
		}
	}

	return count
}

func Max(array [][]int, f func([][]int, int, int) int) int {
	max := -1

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			if fn := f(array, i, j); fn > max {
				max = fn
			}
		}
	}

	return max
}

func Reduce[T, M any](array []T, f func(M, T) M, initialValue M) M {
	accumulator := initialValue

	for _, n := range array {
		accumulator = f(accumulator, n)
	}

	return accumulator
}

func main() {
	handle, err := os.Open("sample.txt")
	check(err)
	defer handle.Close()

	forest := ReadForest(handle)

	isTreeVisible := func(forest [][]int, latitude, longitude int) bool {
		height := forest[latitude][longitude]
		isVisibleFromNorth := func() bool {
			for i := latitude - 1; i >= 0; i-- {
				if forest[i][longitude] >= height {
					return false
				}
			}
			return true
		}

		isVisibleFromEast := func() bool {
			for j := longitude + 1; j < len(forest[0]); j++ {
				if forest[latitude][j] >= height {
					return false
				}
			}
			return true
		}

		isVisibleFromSouth := func() bool {
			for i := latitude + 1; i < len(forest); i++ {
				if forest[i][longitude] >= height {
					return false
				}
			}
			return true
		}

		isVisibleFromWest := func() bool {
			for j := longitude - 1; j >= 0; j-- {
				if forest[latitude][j] >= height {
					return false
				}
			}
			return true
		}

		return isVisibleFromNorth() || isVisibleFromEast() || isVisibleFromSouth() || isVisibleFromWest()
	}

	scenicScore := func(forest [][]int, latitude, longitude int) int {
		northScenicScore := func() int {
			if latitude == 0 {
				return 0
			}
			if forest[latitude-1][longitude] >= forest[latitude][longitude] {
				return 1
			}
			scenicScore := 1
			for i := latitude - 1; i > 0; i-- {
				if forest[i][longitude] >= forest[latitude][longitude] {
					return scenicScore
				}
				if forest[i-1][longitude] <= forest[i][longitude] {
					return scenicScore + 1
				}
				scenicScore++
			}
			return scenicScore
		}

		eastScenicScore := func() int {
			if longitude == len(forest[0])-1 {
				return 0
			}
			if forest[latitude][longitude+1] >= forest[latitude][longitude] {
				return 1
			}
			scenicScore := 1
			for j := longitude + 1; j < len(forest[0])-1; j++ {
				if forest[latitude][j] >= forest[latitude][longitude] {
					return scenicScore
				}
				if forest[latitude][j+1] <= forest[latitude][j] {
					return scenicScore + 1
				}
				scenicScore++
			}
			return scenicScore
		}

		southScenicScore := func() int {
			if latitude == len(forest)-1 {
				return 0
			}
			if forest[latitude+1][longitude] >= forest[latitude][longitude] {
				return 1
			}
			scenicScore := 1
			for i := latitude + 1; i < len(forest)-1; i++ {
				if forest[i][longitude] >= forest[latitude][longitude] {
					return scenicScore
				}
				if forest[i+1][longitude] <= forest[i][longitude] {
					return scenicScore + 1
				}
				scenicScore++
			}
			return scenicScore
		}

		westScenicScore := func() int {
			if longitude == 0 {
				return 0
			}
			if forest[latitude][longitude-1] >= forest[latitude][longitude] {
				return 1
			}
			scenicScore := 1
			for j := longitude - 1; j > 0; j-- {
				if forest[latitude][j] >= forest[latitude][longitude] {
					return scenicScore
				}
				if forest[latitude][j-1] <= forest[latitude][j] {
					return scenicScore + 1
				}
				scenicScore++
			}
			return scenicScore
		}

		if forest[latitude][longitude] == 0 {
			return 0
		}

		return Reduce([]int{northScenicScore(), eastScenicScore(), southScenicScore(), westScenicScore()},
			func(accumulator int, current int) int {
				return accumulator * current
			}, 1)
	}

	fmt.Printf("Part 1: %d\n", Count(forest, isTreeVisible))
	fmt.Printf("Part 2: %d\n", Max(forest, scenicScore))
}
