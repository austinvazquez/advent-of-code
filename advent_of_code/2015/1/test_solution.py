import pytest
import solution

def test_traverse():
    assert solution.traverse("(") == 1
    assert solution.traverse(")") == -1


@pytest.mark.parametrize(
        "input,expected_floor",
        [("(())", 0), ("()()", 0), ("(((", 3), ("(()(()(", 3), ("())", -1), ("))(", -1)]
)
def test_step_1(input, expected_floor):
    def step_1(directions):
        return sum(map(solution.traverse, directions))

    assert step_1(input) == expected_floor


@pytest.mark.parametrize(
        "input,expected_position", 
        [(")", 1), ("())", 3), ("()())", 5)]
)
def test_find_enter_basement(input, expected_position):
    actual = solution.find_enter_basement(input)
    assert actual == expected_position