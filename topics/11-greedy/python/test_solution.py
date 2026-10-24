"""11 贪心的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


def test_activity_selection() -> None:
    start = [1, 3, 0, 5, 8, 5]
    end = [2, 4, 6, 7, 9, 9]
    assert solution.activity_selection(start, end) == [0, 1, 3, 4]
    assert solution.activity_selection([], []) == []


def test_fractional_knapsack() -> None:
    assert solution.fractional_knapsack([10, 20, 30], [60, 100, 120], 50) == 240.0
    assert solution.fractional_knapsack([], [], 10) == 0


@pytest.mark.parametrize(
    ("nums", "expected"),
    [([2, 3, 1, 1, 4], True), ([3, 2, 1, 0, 4], False), ([0], True)],
)
def test_jump_game(nums: list[int], expected: bool) -> None:
    assert solution.jump_game(nums) is expected


@pytest.mark.parametrize(
    ("nums", "expected"),
    [
        ([2, 3, 1, 1, 4], 2),
        ([2, 3, 0, 1, 4], 2),
        ([0], 0),
        ([1, 1, 1, 1], 3),
        ([0, 1], -1),
    ],
)
def test_min_jumps(nums: list[int], expected: int) -> None:
    assert solution.min_jumps(nums) == expected


def test_find_content_children() -> None:
    assert solution.find_content_children([1, 2, 3], [1, 1]) == 1
    assert solution.find_content_children([1, 2], [1, 2, 3]) == 2
