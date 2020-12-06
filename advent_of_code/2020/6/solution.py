"""
Solution for AoC 2020 Day 6
https://adventofcode.com/2020/day/6
"""
from collections import defaultdict
from typing import List, Set


def separate_groups(lines: List[str]) -> List[List[str]]:
    groups = []
    group = []
    line_no = 0

    while line_no < len(lines):
        line = lines[line_no].strip()

        if line:
            group.append(line)
        else:
            groups.append(group)
            group = []

        line_no += 1

    groups.append(group)

    return groups


def questions_anyone_answered(group_answers: List[str]) -> Set:
    answers = set()

    for answer in group_answers:
        list(map(answers.add, answer))

    return answers


def questions_everyone_answered(group_answers: List[str]) -> Set:
    answers = set()

    counter = defaultdict(int)

    for answer in group_answers:
        for ch in answer:
            counter[ch] += 1

    for ch in counter.keys():
        if counter[ch] == len(group_answers):
            answers.add(ch)

    return answers


def main():
    with open("sample.txt") as handle:
        lines = handle.readlines()
    groups = separate_groups(lines)

    count = sum([len(group) for group in list(map(questions_anyone_answered, groups))])
    print(f"Sample: {count}")

    count = sum(
        [len(group) for group in list(map(questions_everyone_answered, groups))]
    )
    print(f"Sample: {count}")

    with open("input.txt") as handle:
        lines = handle.readlines()
    groups = separate_groups(lines)

    count = sum([len(group) for group in list(map(questions_anyone_answered, groups))])
    print(f"Step 1: {count}")

    count = sum([len(group) for group in list(map(questions_everyone_answered, groups))])
    print(f"Step 2: {count}")


if __name__ == "__main__":
    main()
