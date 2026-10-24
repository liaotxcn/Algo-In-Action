"""03 TopK 排行的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_top_k() -> None:
    assert solution.top_k([3, 1, 5, 12, 2, 11], 3) == [12, 11, 5]
    assert solution.top_k([1, 1, 2, 2, 3], 3) == [3, 2, 2]
    assert solution.top_k([], 3) == []
    assert solution.top_k([1, 2], 0) == []


def test_top_k_clamps_k() -> None:
    assert solution.top_k([4, 2, 9], 10) == [9, 4, 2]


def test_top_k_frequent() -> None:
    words = ["i", "love", "leetcode", "i", "love", "coding"]
    assert solution.top_k_frequent(words, 2) == ["i", "love"]


def test_top_k_frequent_tie_break() -> None:
    assert solution.top_k_frequent(["b", "a", "b", "a", "c"], 2) == ["a", "b"]
