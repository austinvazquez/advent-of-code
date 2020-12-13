"""
Solution for AoC 2020 Day 7
https://adventofcode.com/2020/day/7
"""
from re import match


def main():
    with open("input.txt") as handle:
        lines = [line.strip() for line in handle]

    bags = {}

    for line in lines:
        bag, content = line.split(" bags contain ")
        bags[bag] = []

        if content == "no other bags.":
            continue

        for subbag in [bag.strip() for bag in content.split(",")]:
            num, _type = match(r"(\d)+ (.+) bag.*", subbag).groups()
            bags[bag].append((_type, int(num)))

    def contains_shiny_gold_bag(bag):
        return (
            bag == "shiny gold" or 
            any(
                contains_shiny_gold_bag(subbag) 
                for subbag, _ in bags[bag]
            )
        )

    def bags_contained(bag):
        return 1 + sum(
            num_of_bags * bags_contained(subbag) 
            for subbag, num_of_bags in bags[bag]
        )

    count = sum(contains_shiny_gold_bag(bag) for bag in bags) - 1
    print(f"Step 1: {count}")

    count = bags_contained("shiny gold") - 1
    print(f"Step 2: {count}")


if __name__ == "__main__":
    main()
