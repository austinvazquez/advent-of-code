from math import prod
from typing import List


def two_number_sum(nums: List[int], target: int) -> List[int]:
    """
    Complexity analysis
    Time: O(n)
    Space: O(n)
    """
    seen = set()

    result = []

    for num in nums:
        if target - num in seen:
            result = [num, target - num]
            break

        seen.add(num)
    
    return result


def three_number_sum(nums: List[int], target: int) -> List[int]:
    """
    Complexity analysis
    Time: O(n^2)
    Space: O(n)
    """
    nums.sort()

    result = []

    for i in range(len(nums) - 2):
        left, right = i + 1, len(nums) - 1

        while left < right:
            current_sum = nums[left] + nums[i] + nums[right]

            if current_sum < target:
                left += 1
            elif current_sum > target:
                right -= 1
            else:
                result = [nums[left], nums[i], nums[right]]
                break

    return result


def main():
    with open("input.txt") as handle:
        nums = [int(line) for line in handle]

    print(f"Step 1: {prod(two_number_sum(nums, 2020))}")
    print(f"Step 2: {prod(three_number_sum(nums, 2020))}")


if __name__ == "__main__":
    main()
