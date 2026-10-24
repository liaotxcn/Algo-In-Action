"""09 搜索的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


@pytest.mark.parametrize(
    ("target", "expected"),
    [(1, 0), (13, 6), (15, 7), (8, -1), (0, -1), (16, -1)],
)
def test_binary_search(target: int, expected: int) -> None:
    nums = [1, 3, 5, 7, 9, 11, 13, 15]
    assert solution.binary_search(nums, target) == expected


def test_binary_search_empty() -> None:
    assert solution.binary_search([], 1) == -1


@pytest.mark.parametrize(
    ("target", "lower", "upper"),
    [
        (2, 1, 4),
        (1, 0, 1),
        (6, 5, 6),
        (3, 4, 4),
        (0, 0, 0),
        (7, 6, 6),
    ],
)
def test_bounds(target: int, lower: int, upper: int) -> None:
    nums = [1, 2, 2, 2, 4, 6]
    assert solution.lower_bound(nums, target) == lower
    assert solution.upper_bound(nums, target) == upper


def test_search_range() -> None:
    nums = [5, 7, 7, 7, 8, 10]
    assert solution.search_range(nums, 7) == (1, 3)
    assert solution.search_range(nums, 6) == (-1, -1)


def test_interpolation_search() -> None:
    nums = [i * 2 for i in range(100)]
    assert solution.interpolation_search(nums, 84) == 42
    assert solution.interpolation_search(nums, 85) == -1
