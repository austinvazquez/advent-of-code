"""
Solution for AoC 2021 Day 0
https://adventofcode.com/2015/day/1
"""


def main():
    directions = []

    with open("input.txt") as handle:
        for line in handle:
            directions = line

    print(f"Step 1: {sum(map(traverse, directions))}")
    print(f"Step 2: {find_enter_basement(directions)}")


def traverse(direction) -> int:
    return 1 if direction == "(" else -1


def find_enter_basement(directions: []) -> int:
    index = 1
    position = 0

    for direction in directions:
        position += traverse(direction)

        if position == -1:
            return index
        index += 1

    return -1


if __name__ == "__main__":
    main()
