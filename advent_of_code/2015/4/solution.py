"""
Solution for AoC 2020 Day 4
https://adventofcode.com/2015/day/4
"""
from hashlib import md5
from re import compile


def main():
    input = "bgvyzdsv"

    print(f"Step 1: {find_key_to_hash(input, '^00000.*')}")
    print(f"Step 2: {find_key_to_hash(input, '^000000.*')}")


def find_key_to_hash(special_key, pattern) -> int:
    key = 1

    regex = compile(pattern)

    while True:
        hash = md5(str(special_key + str(key)).encode()).hexdigest()

        if regex.match(hash) != None:
            return key

        key += 1


if __name__ == "__main__":
    main()
