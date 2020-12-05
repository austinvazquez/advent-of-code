"""
Solution for AoC 2020 Day 5
https://adventofcode.com/2020/day/5
"""
from dataclasses import dataclass
from math import ceil
from typing import List


SEAT_PER_ROW = 8


@dataclass
class Coordinate:
    row: int
    col: int


def get_seat_id_from(coordinate: Coordinate) -> int:
    return (coordinate.row * SEAT_PER_ROW) + coordinate.col


def binary_partition(partition: str, upper_bounds: int, lower_divisor: str) -> int:
    left, right = 0, upper_bounds
    for ch in partition:
        if ch == lower_divisor:
            right = left + (right - left) // 2
        else:
            left = left + (right - left) // 2

    return ceil((left + right) / 2)


def transform_partition_into_coordinate(partition: str) -> Coordinate:
    return Coordinate(
        row=binary_partition(partition[:7], 127, "F"),
        col=binary_partition(partition[7:], 7, "L"),
    )


def find_seat(seats: List[int]) -> int:
    my_seat = 0

    for seat in seats:
        if seat + 1 not in seats and seat + 2 in seats:
            my_seat = seat + 1
            break

    return my_seat


def main():
    with open("input.txt") as handle:
        boarding_passes = [line.strip() for line in handle]

    boarding_passes = list(
        map(get_seat_id_from, map(transform_partition_into_coordinate, boarding_passes))
    )

    print(f"Step 1: {max(boarding_passes)}")
    print(f"Step 2: {find_seat(boarding_passes)}")


if __name__ == "__main__":
    main()
