"""04 哈希表的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


def test_two_sum() -> None:
    assert solution.two_sum([2, 7, 11, 15], 9) == [0, 1]
    assert solution.two_sum([3, 3], 6) == [0, 1]
    assert solution.two_sum([1, 2], 10) is None


@pytest.mark.parametrize(
    ("nums", "expected"),
    [
        ([100, 4, 200, 1, 3, 2], 4),
        ([0, 3, 7, 2, 5, 8, 4, 6, 0, 1], 9),
        ([], 0),
        ([1, 1, 1], 1),
    ],
)
def test_longest_consecutive(nums: list[int], expected: int) -> None:
    assert solution.longest_consecutive(nums) == expected


def test_group_anagrams() -> None:
    got = solution.group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
    assert got == [["ate", "eat", "tea"], ["bat"], ["nat", "tan"]]
    assert solution.group_anagrams([]) == []


@pytest.mark.parametrize(
    ("nums", "k", "expected"),
    [
        ([1, 1, 1], 2, 2),
        ([1, 2, 3], 3, 2),
        ([1, -1, 0], 0, 3),
    ],
)
def test_subarray_sum(nums: list[int], k: int, expected: int) -> None:
    assert solution.subarray_sum(nums, k) == expected


def test_is_anagram() -> None:
    assert solution.is_anagram("anagram", "nagaram") is True
    assert solution.is_anagram("rat", "car") is False


def test_normalize_word() -> None:
    assert solution.normalize_word("eat") == "aet"
