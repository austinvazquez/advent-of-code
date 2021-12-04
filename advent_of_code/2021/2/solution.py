"""
Solution for AoC 2021 Day 2
https://adventofcode.com/2021/day/2
"""
from collections import namedtuple
from typing import List


Instruction = namedtuple("Instruction", ["direction", "distance"])
Coordinates = namedtuple("Coordinates", ["horizontal", "depth", "aim"])


def execute_instructions(instructions: List[Instruction]) -> Coordinates:
    location = Coordinates(0, 0, 0)

    for instruction in instructions:
        if instruction.direction == "forward":
            location = Coordinates(
                location.horizontal + instruction.distance, location.depth, location.aim
            )
        elif instruction.direction == "up":
            location = Coordinates(
                location.horizontal, location.depth - instruction.distance, location.aim
            )
        else:  # instruction.direction == "down"
            location = Coordinates(
                location.horizontal, location.depth + instruction.distance, location.aim
            )

    return location


def execute_instructions_with_aim(instructions: List[Instruction]) -> Coordinates:
    location = Coordinates(0, 0, 0)

    for instruction in instructions:
        if instruction.direction == "forward":
            location = Coordinates(
                location.horizontal + instruction.distance,
                location.depth + (instruction.distance * location.aim),
                location.aim,
            )
        elif instruction.direction == "up":
            location = Coordinates(
                location.horizontal, location.depth, location.aim - instruction.distance
            )
        else:  # instruction.direction == "down"
            location = Coordinates(
                location.horizontal, location.depth, location.aim + instruction.distance
            )

    return location


def multiply_coordinates(coordinates: Coordinates) -> int:
    return coordinates.depth * coordinates.horizontal


def main():
    def parse_instruction_from_text(line: str) -> Instruction:
        direction, distance = line.split(" ")
        return Instruction(direction, int(distance))

    with open("input.txt") as handle:
        instructions = [parse_instruction_from_text(line) for line in handle]

    print(f"Step 1: {multiply_coordinates(execute_instructions(instructions))}")
    print(
        f"Step 2: {multiply_coordinates(execute_instructions_with_aim(instructions))}"
    )


if __name__ == "__main__":
    main()
