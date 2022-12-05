package main

// Solution for AoC 2021 Day 2
// https://adventofcode.com/2021/day/2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type instruction struct {
	direction string
	distance  int
}

func ReadInstructions(r io.Reader) ([]instruction, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var instructions []instruction

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		var instruction instruction
		instruction.direction = split[0]
		val, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		instruction.distance = val
		instructions = append(instructions, instruction)
	}

	return instructions, scanner.Err()
}

type coordinate struct {
	horizontal int
	depth      int
	aim        int
}

func ExecuteInstructions(instructions []instruction) coordinate {
	location := coordinate{0, 0, 0}

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			location.horizontal += instruction.distance
		case "up":
			location.depth -= instruction.distance
		case "down":
			location.depth += instruction.distance
		}
	}

	return location
}

func ExecuteInstructionsWithAim(instructions []instruction) coordinate {
	location := coordinate{0, 0, 0}

	for _, instruction := range instructions {
		switch instruction.direction {
		case "forward":
			location.horizontal += instruction.distance
			location.depth += (location.aim * instruction.distance)
		case "up":
			location.aim -= instruction.distance
		case "down":
			location.aim += instruction.distance
		}
	}

	return location
}

func MultiplyCoordinates(coordinate coordinate) int {
	return coordinate.horizontal * coordinate.depth
}

func main() {
	handle, err := os.Open("input.txt")
	check(err)

	defer handle.Close()

	instructions, err := ReadInstructions(handle)
	check(err)

	fmt.Printf("Part 1: %d\n", MultiplyCoordinates(ExecuteInstructions(instructions)))
	fmt.Printf("Part 2: %d\n", MultiplyCoordinates(ExecuteInstructionsWithAim(instructions)))
}
