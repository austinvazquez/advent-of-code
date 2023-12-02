"""
Solution for AoC 2015 Day 3
https://adventofcode.com/2015/day/3
"""


def main():
    directions = []

    with open("input.txt") as handle:
        directions = handle.read()

    print(f"Step 1: {compute_number_of_houses(directions)}")
    print(f"Step 2: {compute_number_of_houses(directions, increment = 2)}")


def compute_number_of_houses(directions, increment=1) -> int:
    path = [(0, 0), (0, 0)]

    for step in directions:
        last_house = path[-increment]

        if step == "^":
            path.append((last_house[0], last_house[1] + 1))
        elif step == ">":
            path.append((last_house[0] + 1, last_house[1]))
        elif step == "<":
            path.append((last_house[0] - 1, last_house[1]))
        else:
            path.append((last_house[0], last_house[1] - 1))

    return len(set(path))


if __name__ == "__main__":
    main()
