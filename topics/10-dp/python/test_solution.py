"""10 动态规划的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


def test_fib() -> None:
    assert [solution.fib(n) for n in range(9)] == [0, 1, 1, 2, 3, 5, 8, 13, 21]
    assert solution.fib(-1) == 0


@pytest.mark.parametrize(("n", "expected"), [(1, 1), (2, 2), (3, 3), (4, 5), (5, 8)])
def test_climb_stairs(n: int, expected: int) -> None:
    assert solution.climb_stairs(n) == expected


def test_knapsack_01() -> None:
    weights = [1, 3, 4]
    values = [15, 20, 30]
    assert solution.knapsack_01(weights, values, 4) == 35
    assert solution.knapsack_01(weights, values, 0) == 0


@pytest.mark.parametrize(
    ("nums", "expected"),
    [
        ([10, 9, 2, 5, 3, 7, 101, 18], 4),
        ([0, 1, 0, 3, 2, 3], 4),
        ([7, 7, 7], 1),
        ([], 0),
    ],
)
def test_lis(nums: list[int], expected: int) -> None:
    assert solution.lis(nums) == expected


@pytest.mark.parametrize(
    ("a", "b", "expected"),
    [("horse", "ros", 3), ("intention", "execution", 5), ("", "abc", 3), ("same", "same", 0)],
)
def test_edit_distance(a: str, b: str, expected: int) -> None:
    assert solution.edit_distance(a, b) == expected


@pytest.mark.parametrize(
    ("coins", "amount", "expected"),
    [([1, 2, 5], 11, 3), ([2], 3, -1), ([1], 0, 0), ([1, 2, 5], 0, 0)],
)
def test_coin_change(coins: list[int], amount: int, expected: int) -> None:
    assert solution.coin_change(coins, amount) == expected
