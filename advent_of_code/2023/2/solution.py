"""
Solution for AoC 2023 Day 2
https://adventofcode.com/2023/day/2
"""
from math import prod


def main():
    with open("input.txt") as file:
        games = [line for line in file]

    print(f"Step 1: {sum(map(is_game_possible, games))}")
    print(f"Step 2: {sum(map(minimum_power_of_game, games))}")


def is_game_possible(game: str) -> int:
    game = game.lstrip("Game ")

    prefix, suffix = game.split(":")

    game_id = int(prefix)
    cube_limits = {
        "red": 12,
        "blue": 14,
        "green": 13,
    }

    for round in suffix.split(";"):
        for color in round.split(","):
            count, color = color.strip().split(" ")
            if int(count) > cube_limits[color]:
                return 0

    return game_id


def minimum_power_of_game(game: str) -> int:
    _, game = game.split(":")

    min_count = {"red": 0, "blue": 0, "green": 0}

    for round in game.split(";"):
        for color in round.split(","):
            count, color = color.strip().split(" ")
            min_count[color] = max(min_count[color], int(count))

    return prod(min_count.values())


if __name__ == "__main__":
    main()
