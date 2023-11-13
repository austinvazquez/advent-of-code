"""
Solution for AoC 2020 Day 5
https://adventofcode.com/2015/day/5
"""
from re import compile


def main():
    christmas_list = []

    with open("input.txt") as handle:
        christmas_list = [line for line in handle]

    assert len(christmas_list) > 0

    print(f"Step 1: {len(list(filter(is_nice_strict, christmas_list)))}")
    print(f"Step 2: {len(list(filter(is_nice_leniant, christmas_list)))}")


def is_nice_strict(string: str) -> bool:
    def contains_three_vowels(string: str) -> bool:
        vowel_count = 0

        for v in "aeiou":
            vowel_count += string.count(v)
        
        return vowel_count >= 3

    def contains_double_letter(string: str) -> bool:
        for i in range(len(string) - 1):
            if string[i] == string[i+1]:
                return True

        return False
    
    def does_not_contain_excluded_strings(string: str) -> bool:
        excluded_strings = ["ab", "cd", "pq", "xy"]

        for exstring in excluded_strings:
            if string.count(exstring) > 0:
                return False

        return True

    return contains_three_vowels(string) and contains_double_letter(string) and does_not_contain_excluded_strings(string)


def is_nice_leniant(string: str) -> bool:
    pair_regex = compile(r"([a-z]{2}).*\1")
    def contains_pair_letters(string: str) -> bool:
        return pair_regex.search(string) != None

    xox_regex = compile(r"([a-z]{1})([a-z]{1})\1")
    def contains_xox(string: str) -> bool:
        return xox_regex.search(string) != None

    return contains_pair_letters(string) and contains_xox(string)


if __name__ == "__main__":
    main()
