"""06 堆的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


def test_max_heap() -> None:
    heap = solution.MaxHeap()
    assert heap.pop() is None
    for value in [5, 2, 8, 1, 9, 3]:
        heap.push(value)
    assert heap.peek() == 9
    assert [heap.pop() for _ in range(6)] == [9, 8, 5, 3, 2, 1]
    assert heap.pop() is None


@pytest.mark.parametrize(
    ("nums", "expected"),
    [
        ([], []),
        ([7], [7]),
        ([5, 4, 3, 2, 1], [1, 2, 3, 4, 5]),
        ([4, 2, 2, 8, 3, 3, 1], [1, 2, 2, 3, 3, 4, 8]),
    ],
)
def test_heap_sort(nums: list[int], expected: list[int]) -> None:
    original = nums.copy()
    assert solution.heap_sort(nums) == expected
    assert nums == original  # 不修改入参


@pytest.mark.parametrize(
    ("nums", "k", "expected"),
    [
        ([3, 2, 1, 5, 6, 4], 2, [6, 5]),
        ([3, 2, 3, 1, 2, 4, 5, 5, 6], 4, [6, 5, 5, 4]),
        ([1, 2], 5, [2, 1]),
        ([1, 2], 0, []),
        ([], 2, []),
    ],
)
def test_top_k(nums: list[int], k: int, expected: list[int]) -> None:
    assert solution.top_k(nums, k) == expected
