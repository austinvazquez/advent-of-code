import solution
import pytest


@pytest.mark.parametrize(
        "input,expected_number_of_houses",
        [(">", 2), ("^>v<", 4), ("^v^v^v^v^v", 2)]
)
def test_compute_number_of_houses(input, expected_number_of_houses):
    assert solution.compute_number_of_houses(input) == expected_number_of_houses

@pytest.mark.parametrize(
        "input,expected_number_of_houses",
        [("^v", 3), ("^>v<", 3), ("^v^v^v^v^v", 11)]
)
def test_compute_number_of_houses_increment_2(input, expected_number_of_houses):
    assert solution.compute_number_of_houses(input, increment=2) == expected_number_of_houses