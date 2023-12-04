"""
Solution for AoC 2023 Day 4
https://adventofcode.com/2023/day/4
"""
from typing import List


def main():
    with open("input.txt") as file:
        cards = [line for line in file]

    print(f"Step 1: {sum(map(lambda c: ScratchCard(c).value(), cards))}")
    print(f"Step 2: {compute_final_scratchcard_count(cards)}")


class ScratchCard:
    def __init__(self, card: str):
        id, game = card.split(":")
        id = id.lstrip("Card ")
        self.id = int(id)

        game = game.strip().split("|")
        self.winning_numbers = set([int(n) for n in game[0].strip().split(" ") if n != ""])
        self.played_numbers = set([int(n) for n in game[1].strip().split(" ") if n != ""])

    def value(self) -> int:
        value = 0
        for n in self.played_numbers:
            if n in self.winning_numbers:
                value = 2 * value if value != 0 else 1
        return value


def compute_final_scratchcard_count(cards: List[str]) -> int:
    return 0

if __name__ == "__main__":
    main()
