"""
Solution for AoC 2021 Day 1
https://adventofcode.com/2021/day/1
"""
from typing import List


def count_increases_with_sliding_window(nums: List[int], window_size: int = 1) -> int:
    """
    Complexity analysis
    Time: O(n)
    Space: O(1)
    """
    count = 0

    for i in range(0, len(nums) - window_size):
        if nums[i + window_size] > nums[i]:
            count += 1

    return count


def main():
    with open("input.txt") as handle:
        nums = [int(line) for line in handle]

    print(f"Step 1: {count_increases_with_sliding_window(nums)}")
    print(f"Step 2: {count_increases_with_sliding_window(nums, sliding_window_size=3)}")


if __name__ == "__main__":
    main()
