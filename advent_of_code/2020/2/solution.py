"""
Solution for AoC 2020 Day 2
https://adventofcode.com/2020/day/2
"""
from typing import Dict, List


def parse_password_string(password: str) -> Dict:
    strings = password.split()

    range_string = strings[0]
    start, end = range_string.split("-")

    return {
        "start": int(start),
        "end": int(end),
        "character": strings[1][0],
        "password": strings[2],
    }


def by_character_count_validation(password: str) -> bool:
    """
    Check the occurrence count of the character in the password policy is within range.
    """
    values = parse_password_string(password)
    start, end = values["start"], values["end"]
    character = values["character"]
    password = values["password"]
    return password.count(character) in range(start, end + 1)


def by_character_position_validation(password: str) -> bool:
    """
    Check the character occurs in a single position from the password policy.
    """
    only_one_occurrence = lambda ch, s1, s2: s1 != s2 and ch in (s1, s2)

    values = parse_password_string(password)
    first, second = values["start"] - 1, values["end"] - 1
    character = values["character"]
    password = values["password"]

    return only_one_occurrence(character, password[first], password[second])


def count_valid_passwords(passwords: List[str], is_valid_by_policy) -> int:
    num_valid_passwords = 0

    for password in passwords:
        if is_valid_by_policy(password):
            num_valid_passwords += 1

    return num_valid_passwords


def main():
    with open("input.txt") as handle:
        passwords = [line for line in handle]

    print(f"Step 1: {count_valid_passwords(passwords, by_character_count_validation)}")
    print(
        f"Step 2: {count_valid_passwords(passwords, by_character_position_validation)}"
    )


if __name__ == "__main__":
    main()
