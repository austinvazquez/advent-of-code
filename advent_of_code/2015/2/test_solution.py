import pytest
import solution


@pytest.mark.parametrize(
    "input,expected_surface_area", [([2, 3, 4], 58), ([1, 1, 10], 43)]
)
def test_compute_surface_area(input, expected_surface_area):
    assert solution.compute_surface_area(input) == expected_surface_area


@pytest.mark.parametrize("input,expected_length", [([2, 3, 4], 34), ([1, 1, 10], 14)])
def test_compute_ribbon_length(input, expected_length):
    assert solution.compute_ribbon_length(input) == expected_length
