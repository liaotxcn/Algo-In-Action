"""03 栈与队列的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("()", True),
        ("()[]{}", True),
        ("(]", False),
        ("([)]", False),
        ("{[]}", True),
        ("", True),
        ("(", False),
    ],
)
def test_is_valid(text: str, expected: bool) -> None:
    assert solution.is_valid(text) is expected


def test_min_stack() -> None:
    stack = solution.MinStack()
    stack.push(-2)
    stack.push(0)
    stack.push(-3)
    assert stack.get_min() == -3
    assert stack.pop() == -3
    assert stack.top() == 0
    assert stack.get_min() == -2


def test_my_queue() -> None:
    queue = solution.MyQueue()
    assert queue.empty() is True
    queue.push(1)
    queue.push(2)
    assert queue.peek() == 1
    assert queue.pop() == 1
    assert queue.empty() is False
    queue.push(3)
    assert queue.pop() == 2
    assert queue.pop() == 3


@pytest.mark.parametrize(
    ("nums", "k", "expected"),
    [
        ([1, 3, -1, -3, 5, 3, 6, 7], 3, [3, 3, 5, 5, 6, 7]),
        ([1], 1, [1]),
        ([9, 8, 7], 5, [9]),
        ([], 3, []),
    ],
)
def test_max_sliding_window(nums: list[int], k: int, expected: list[int]) -> None:
    assert solution.max_sliding_window(nums, k) == expected
