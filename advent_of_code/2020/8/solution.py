"""
Solution for AoC 2020 Day 8
https://adventofcode.com/2020/day/8
"""
from copy import deepcopy
from typing import List


def run(instructions: List) -> int:
    seen = set()
    accumulator = i = 0

    while i < len(instructions) and i not in seen:
        seen.add(i)
        instruction = instructions[i]

        if instruction[0] == "acc":
            accumulator += instruction[1]
        elif instruction[0] == "jmp":
            i += instruction[1] - 1

        i += 1

    return i == len(instructions), accumulator


def main():
    with open("input.txt") as handle:
        instructions = handle.readlines()

    def split_instruction(instruction):
        instruction, value = instruction.split()
        return [instruction, int(value)]

    instructions = list(map(split_instruction, instructions))

    _, accumulator = run(instructions)
    print(f"Step 1: {accumulator}")

    for i, instruction in enumerate(instructions):
        if instruction[0] == "acc":
            continue

        new_set = deepcopy(instructions)
        new_set[i][0] = "jmp" if new_set[i][0] == "nop" else "nop"
        terminated, accumulator = run(new_set)

        if terminated:
            break

    print(f"Step 2: {accumulator}")


if __name__ == "__main__":
    main()
