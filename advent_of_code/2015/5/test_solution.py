import pytest
import solution


@pytest.mark.parametrize(
    "input,is_nice",
    [
        ("jchzalrnumimnmhp", False),
        ("haegwjzuvuyypxyu", False),
        ("dvszwmarrgswjxmb", False),
        ("ugknbfddgicrmopn", True),
    ],
)
def test_is_nice_strict(input, is_nice):
    assert solution.is_nice_strict(input) == is_nice


@pytest.mark.parametrize(
    "input,is_nice",
    [
        ("ieodomkazucvgmuy", False),
        ("uurcxstgmygtbstg", False),
        ("qjhvhtzxzqqjkmpb", True),
        ("xxyxx", True),
    ],
)
def test_is_nice_leniant(input, is_nice):
    assert solution.is_nice_leniant(input) == is_nice
