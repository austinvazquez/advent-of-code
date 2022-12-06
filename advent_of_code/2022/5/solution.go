package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Solution for AoC 2022 Day 5
// https://adventofcode.com/2022/day/5

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewStacks(textLines []string) [][]rune {
	reverse := func(array []string) []string {
		rarray := make([]string, len(array))
		for i, j := len(array)-1, 0; i >= 0; i, j = i-1, j+1 {
			rarray[j] = array[i]
		}
		return rarray
	}

	textLines = reverse(textLines)

	stackLine := strings.Fields(textLines[0])
	numberOfStacks, err := strconv.Atoi(stackLine[len(stackLine)-1])
	check(err)

	stacks := make([][]rune, numberOfStacks)

	for i := 1; i < len(textLines); i++ {
		for j, k := 1, 0; j < len(textLines[i]); j, k = j+4, k+1 {
			toAppend := []rune(textLines[i])[j]
			if string(toAppend) != " " {
				stacks[k] = append(stacks[k], toAppend)
			}
		}
	}

	return stacks
}

func Apply9000(stacks [][]rune, moves []*Move) {
	for _, m := range moves {
		for i := 0; i < m.numberOfCrates; i++ {
			m.DoOnce(stacks)
		}
	}
}

func Apply9001(stacks [][]rune, moves []*Move) {
	for _, m := range moves {
		m.Do(stacks)
	}
}

func ToString(stacks [][]rune) string {
	var result []rune
	for _, stack := range stacks {
		if len(stack) > 0 {
			result = append(result, stack[len(stack)-1])
		} else {
			result = append(result, rune(20))
		}
	}
	return string(result)
}

type Move struct {
	numberOfCrates int
	from           int
	to             int
}

func NewMove(move string) *Move {
	words := strings.Split(move, " ")
	numberOfCrates, err := strconv.Atoi(words[1])
	check(err)

	from, err := strconv.Atoi(words[3])
	check(err)

	to, err := strconv.Atoi(words[5])
	check(err)

	return &Move{
		numberOfCrates: numberOfCrates,
		from:           from,
		to:             to,
	}
}

func (m *Move) DoOnce(stacks [][]rune) {
	before := m.numberOfCrates
	defer func() {
		m.numberOfCrates = before
	}()
	m.numberOfCrates = 1
	m.Do(stacks)
}

func (m *Move) Do(stacks [][]rune) {
	prepend := func(array []rune, r rune) []rune {
		result := []rune{r}
		return append(result, array...)
	}

	from, to := m.from-1, m.to-1
	movedCount := 0

	var runesToMove []rune
	for i := 0; i < m.numberOfCrates; i++ {
		if len(stacks[from])-movedCount == 0 {
			break
		}
		runesToMove = prepend(runesToMove, stacks[from][len(stacks[from])-movedCount-1])
		movedCount++
	}

	for _, r := range runesToMove {
		stacks[to] = append(stacks[to], r)
	}

	if movedCount == len(stacks[from]) {
		stacks[from] = []rune{}
	} else {
		stacks[from] = stacks[from][0 : len(stacks[from])-movedCount]
	}
}

func ReadStacksAndMoves(r io.Reader) ([][]rune, []*Move) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var stackLines []string
	var moves []*Move

	var scannedStacks bool
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if scannedStacks {
			moves = append(moves, NewMove(line))
		} else if strings.Contains(line, "1") {
			stackLines = append(stackLines, line)
			scannedStacks = true
		} else {
			stackLines = append(stackLines, line)
		}
	}

	return NewStacks(stackLines), moves
}

func DeepCopy(stacks [][]rune) [][]rune {
	copy := make([][]rune, len(stacks))
	for i, stack := range stacks {
		crates := make([]rune, len(stack))
		for j, crate := range stack {
			crates[j] = crate
		}
		copy[i] = crates
	}
	return copy
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	stacks, moves := ReadStacksAndMoves(handle)

	copy := DeepCopy(stacks)
	Apply9000(copy, moves)

	fmt.Printf("Part 1: %s\n", ToString(copy))

	copy = DeepCopy(stacks)
	Apply9001(copy, moves)

	fmt.Printf("Part 2: %s\n", ToString(copy))
}
