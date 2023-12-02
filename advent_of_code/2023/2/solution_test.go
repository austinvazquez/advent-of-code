package main

import "testing"

func TestIsGamePossible(t *testing.T) {
	testCases := []struct {
		Game     string
		Expected int
	}{
		{
			Game:     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: 1,
		},
		{
			Game:     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: 2,
		},
		{
			Game:     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: 0,
		},
		{
			Game:     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: 0,
		},
		{
			Game:     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Game, func(t *testing.T) {
			actual := IsGamePossible(testCase.Game)
			if actual != testCase.Expected {
				t.Fatalf("Expected %d, got %d", testCase.Expected, actual)
			}
		})
	}
}

func TestMinimumPowerOfGame(t *testing.T) {
	testCases := []struct {
		Game     string
		Expected int
	}{
		{
			Game:     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: 48,
		},
		{
			Game:     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: 12,
		},
		{
			Game:     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: 1560,
		},
		{
			Game:     "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: 630,
		},
		{
			Game:     "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: 36,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Game, func(t *testing.T) {
			actual := MinimumPowerOfGame(testCase.Game)
			if actual != testCase.Expected {
				t.Fatalf("Expected %d, got %d", testCase.Expected, actual)
			}
		})
	}
}
