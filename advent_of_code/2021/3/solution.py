"""
Solution for AoC 2021 Day 3
https://adventofcode.com/2021/day/3
"""
from typing import List, Tuple


def get_gamma_and_epsilon_rates(numbers: List[str]) -> Tuple[int, int]:
    gamma = "".join(
        [
            "1" if column.count("1") >= column.count("0") else "0"
            for column in transpose(numbers)
        ]
    )
    epsilon = inverse(gamma)

    return int(gamma, 2), int(epsilon, 2)


def filter_numbers(numbers: List[str], inverse: bool = False) -> List[str]:
    for i in range(len(numbers)):
        transposed_numbers = transpose(numbers)
        col = transposed_numbers[i]

        if inverse:
            numbers = list(
                filter(
                    lambda n: n[i] == "0" if col.count("1") >= col.count("0") else n[i] == "1",
                    numbers,
                )
            )
        else:
            numbers = list(
                filter(
                    lambda n: n[i] == "1" if col.count("1") >= col.count("0") else n[i] == "0",
                    numbers,
                )
            )

        if len(numbers) == 1:
            break

    return numbers[0]


def get_life_support_rates(numbers: List[str]) -> Tuple[int, int]:
    oxygen_generator = filter_numbers(numbers)
    carbon_dioxide_scrubber = filter_numbers(numbers, inverse=True)

    return int(oxygen_generator, 2), int(carbon_dioxide_scrubber, 2)


def transpose(numbers: List[str]) -> List[str]:
    return ["".join(row[i] for row in numbers) for i in range(len(numbers[0]))]


def inverse(numbers: str) -> str:
    return "".join(map(lambda bit: "1" if bit == "0" else "0", numbers))


def main():
    with open("input.txt") as handle:
        numbers = [line.strip() for line in handle]

    gamma_rate, epsilon_rate = get_gamma_and_epsilon_rates(numbers)
    power_consumption_rate = gamma_rate * epsilon_rate

    print(f"Part 1: {power_consumption_rate}")

    oxygen_generator_rate, carbon_dioxide_scrubber_rate = get_life_support_rates(
        numbers
    )
    life_support_rate = oxygen_generator_rate * carbon_dioxide_scrubber_rate

    print(f"Part 2: {life_support_rate}")


if __name__ == "__main__":
    main()
