"""
Solution for AoC 2020 Day 4
https://adventofcode.com/2020/day/4
"""
from typing import Dict, List


def parse_passports_from_input(
    lines: List[str], required_fields: List[str]
) -> List[Dict]:
    passport = {}
    passports = []

    for line in lines:
        line = line.strip()

        if line:
            for pair in line.split(" "):
                key, value = pair.split(":")
                passport[key] = value
        else:
            missing_fields = [
                field for field in required_fields if field not in passport.keys()
            ]

            if not missing_fields:
                passports.append(passport)

            passport = {}

    return passports


def filter_invalid_birth_years(passport) -> bool:
    return 1920 <= int(passport["byr"]) <= 2002


def filter_invalid_issue_years(passport: Dict) -> bool:
    return 2010 <= int(passport["iyr"]) <= 2020


def filter_invalid_expiration_years(passport: Dict) -> bool:
    return 2020 <= int(passport["eyr"]) <= 2030


def filter_invalid_heights(passport: Dict) -> bool:
    if "cm" in passport["hgt"]:
        return 150 <= int(passport["hgt"].split("c")[0]) <= 193
    elif "in" in passport["hgt"]:
        return 59 <= int(passport["hgt"].split("i")[0]) <= 76
    else:
        return False


def filter_invalid_hair_colors(passport: Dict) -> bool:
    return (
        len(passport["hcl"]) == 7
        and passport["hcl"][0] == "#"
        and all([ch in ("0123456789abcdef") for ch in passport["hcl"][1:]])
    )


def filter_invalid_eye_colors(passport: Dict) -> bool:
    return passport["ecl"] in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]


def filter_invalid_pids(passport: Dict) -> bool:
    return len(passport["pid"]) == 9 and all(
        [digit.isdigit() for digit in passport["pid"]]
    )


def main():
    with open("sample.txt") as handle:
        lines = [line for line in handle]

    fields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"]

    passports = parse_passports_from_input(lines, fields)

    print(f"Sample: {len(passports)}")

    with open("input.txt") as handle:
        lines = [line for line in handle]

    passports = parse_passports_from_input(lines, fields)
    print(f"Step 1: {len(passports)}")

    filters = (
        filter_invalid_birth_years,
        filter_invalid_issue_years,
        filter_invalid_expiration_years,
        filter_invalid_heights,
        filter_invalid_hair_colors,
        filter_invalid_eye_colors,
        filter_invalid_pids,
    )
    passports = list(
        filter(lambda passport: all(f(passport) for f in filters), passports)
    )
    print(f"Step 2: {len(passports)}")


if __name__ == "__main__":
    main()
