"""
Solution for AoC 2020 Day 3
https://adventofcode.com/2020/day/3
"""
from math import prod
from typing import List


def count_trees_in_path(grid: List[List[str]], pathing: List[int]) -> int:
    """
    Complexity analysis
    Time: O(n) - where n is the number of rows in grid.
    Space: O(1) - constant spacing required.
    """
    num_trees_in_path = 0
    row, col = 0, 0

    while row < len(grid):
        col += pathing[0]
        row += pathing[1]

        if row >= len(grid):
            break

        if col >= len(grid[row]):
            col %= len(grid[row])

        if grid[row][col] == "#":
            num_trees_in_path += 1

    return num_trees_in_path


def main():
    with open("sample.txt") as handle:
        sample_grid = [list(line.strip()) for line in handle]

    print(f"Sample: {count_trees_in_path(sample_grid, pathing=[3, 1])}")

    with open("input.txt") as handle:
        grid = [list(line.strip()) for line in handle]

    print(f"Step 1: {count_trees_in_path(grid, pathing=[3, 1])}")

    encounters = [
        count_trees_in_path(grid, pathing)
        for pathing in [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]
    ]
    print(f"Step 2: {prod(encounters)}")


if __name__ == "__main__":
    main()
