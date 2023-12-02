"""
Solution for AoC 2015 Day 2
https://adventofcode.com/2015/day/2
"""


def main():
    boxes = []

    with open("input.txt") as handle:
        boxes = [
            list(map(lambda ch: int(ch), line.strip().split("x"))) for line in handle
        ]

    print(f"Step 1: {sum(map(compute_surface_area, boxes))}")
    print(f"Step 2: {sum(map(compute_ribbon_length, boxes))}")


def compute_surface_area(size):
    length, width, height = size
    side_1, side_2, side_3 = length * width, width * height, height * length

    return (2 * side_1) + (2 * side_2) + (2 * side_3) + min(side_1, side_2, side_3)


def compute_ribbon_length(size):
    length, width, height = size
    perimeter_1 = (2 * length) + (2 * width)
    perimeter_2 = (2 * width) + (2 * height)
    perimeter_3 = (2 * height) + (2 * length)

    return min(perimeter_1, perimeter_2, perimeter_3) + (length * width * height)


if __name__ == "__main__":
    main()
