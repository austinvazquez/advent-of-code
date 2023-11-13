import pytest
import solution

@pytest.mark.parametrize(
    "key,pattern,expected",
    [("abcdef", "^00000.*", 609043),
     ("pqrstuv", "^00000.*", 1048970)]
)
def test_find_key_to_hash(key, pattern, expected):
    assert solution.find_key_to_hash(key, pattern) == expected