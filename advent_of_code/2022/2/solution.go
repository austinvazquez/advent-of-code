package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Solution for AoC 2022 Day 2
// https://adventofcode.com/2022/day/2

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3

	ShouldLose = 1
	ShouldDraw = 2
	ShouldWin  = 3
)

type Sign struct {
	value int
}

func NewSign(character string) (*Sign, error) {
	switch character {
	case "A":
		return &Sign{Rock}, nil
	case "B":
		return &Sign{Paper}, nil
	case "C":
		return &Sign{Scissors}, nil
	case "X":
		return &Sign{ShouldLose}, nil
	case "Y":
		return &Sign{ShouldDraw}, nil
	case "Z":
		return &Sign{ShouldWin}, nil
	default:
		return nil, fmt.Errorf("invalid sign %s", character)
	}
}

func (s *Sign) Value() int {
	return s.value
}

func (s *Sign) Beats(other *Sign) bool {
	switch s.value {
	case Rock:
		return other.value == Scissors
	case Paper:
		return other.value == Rock
	case Scissors:
		return other.value == Paper
	default:
		panic("Invalid sign")
	}
}

func (s *Sign) WinningMove() int {
	switch s.value {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("Invalid sign")
	}
}

func (s *Sign) LosingMove() int {
	switch s.value {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		panic("Invalid sign")
	}
}

type Round struct {
	expected  *Sign
	suggested *Sign
}

func (r *Round) ScoreFromSuggestedMove() int {
	extra := 0
	if r.expected.Value() == r.suggested.Value() {
		extra = 3
	} else if r.suggested.Beats(r.expected) {
		extra = 6
	}
	return r.suggested.Value() + extra
}

func (r *Round) ScoreFromSuggestedOutcome() int {
	if r.suggested.Value() == ShouldDraw {
		return 3 + r.expected.Value()
	} else if r.suggested.Value() == ShouldWin {
		return 6 + r.expected.WinningMove()
	}
	return r.expected.LosingMove()
}

func ReadRounds(r io.Reader) ([]Round, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var rounds []Round
	for scanner.Scan() {
		expected, err := NewSign(scanner.Text())
		if err != nil {
			return rounds, err
		}
		scanner.Scan()
		suggested, err := NewSign(scanner.Text())
		if err != nil {
			return rounds, err
		}
		rounds = append(rounds, Round{expected, suggested})
	}

	return rounds, nil
}

func Reduce(array []Round, f func(Round) int) int {
	accumulator := 0

	for _, n := range array {
		accumulator += f(n)
	}

	return accumulator
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)
	defer handle.Close()

	rounds, err := ReadRounds(handle)
	check(err)

	scoreFromSuggestedMove := func(r Round) int {
		return r.ScoreFromSuggestedMove()
	}

	scoreFromSuggestedOutcome := func(r Round) int {
		return r.ScoreFromSuggestedOutcome()
	}

	fmt.Printf("Part 1: %d\n", Reduce(rounds, scoreFromSuggestedMove))
	fmt.Printf("Part 2: %d\n", Reduce(rounds, scoreFromSuggestedOutcome))
}
