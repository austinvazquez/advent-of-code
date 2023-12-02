package main

// Solution for AoC 2023 Day 2
// https://adventofcode.com/2023/day/2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer file.Close()

	games := ReadGames(file)

	sum := func(l, r int) int {
		return l + r
	}

	fmt.Printf("Step 1: %d\n", Reduce(Map(games, IsGamePossible), sum, 0))
	fmt.Printf("Step 2: %d\n", Reduce(Map(games, MinimumPowerOfGame), sum, 0))
}

func ReadGames(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var games []string

	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	return games
}

func Reduce[T any](collection []T, f func(T, T) T, initial T) T {
	accumulator := initial

	for _, c := range collection {
		accumulator = f(accumulator, c)
	}

	return accumulator
}

func Map[T, M any](collection []T, f func(T) M) []M {
	fcs := []M{}

	for _, c := range collection {
		fcs = append(fcs, f(c))
	}

	return fcs
}

func IsGamePossible(game string) int {
	game = strings.TrimPrefix(game, "Game ")

	prefix, suffix, ok := strings.Cut(game, ":")
	if !ok {
		log.Fatalf("err: could not split game: %s", game)
	}

	id, err := strconv.Atoi(prefix)
	if err != nil {
		log.Fatalf("err: could not cast game id: %v", err)
	}

	cubeLimit := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	for _, round := range strings.Split(suffix, ";") {
		for _, color := range strings.Split(round, ",") {
			color = strings.TrimPrefix(color, " ")

			countS, color, ok := strings.Cut(color, " ")
			if !ok {
				log.Fatalf("err: could not cut round: %s", round)
			}

			count, err := strconv.Atoi(countS)
			if err != nil {
				log.Fatalf("err: cound not cast count: %v", err)
			}

			if count > cubeLimit[color] {
				return 0
			}
		}
	}

	return id
}

func MinimumPowerOfGame(game string) int {
	_, suffix, ok := strings.Cut(game, ":")
	if !ok {
		log.Fatalf("err: could not split game: %s", game)
	}

	minCubeCount := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, round := range strings.Split(suffix, ";") {
		for _, color := range strings.Split(round, ",") {
			color = strings.TrimPrefix(color, " ")

			countS, color, ok := strings.Cut(color, " ")
			if !ok {
				log.Fatalf("err: could not cut round: %s", round)
			}

			count, err := strconv.Atoi(countS)
			if err != nil {
				log.Fatalf("err: cound not cast count: %v", err)
			}

			if count > minCubeCount[color] {
				minCubeCount[color] = count
			}
		}
	}

	return minCubeCount["red"] * minCubeCount["blue"] * minCubeCount["green"]
}
