"""12 回溯的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


def test_permutations() -> None:
    assert solution.permutations([1, 2, 3]) == [
        [1, 2, 3],
        [1, 3, 2],
        [2, 1, 3],
        [2, 3, 1],
        [3, 1, 2],
        [3, 2, 1],
    ]
    assert solution.permutations([1]) == [[1]]


def test_permutations_unique() -> None:
    assert solution.permutations_unique([1, 1, 2]) == [
        [1, 1, 2],
        [1, 2, 1],
        [2, 1, 1],
    ]


def test_combine() -> None:
    assert solution.combine(4, 2) == [
        [1, 2],
        [1, 3],
        [1, 4],
        [2, 3],
        [2, 4],
        [3, 4],
    ]
    assert solution.combine(4, 0) == []


def test_subsets() -> None:
    assert solution.subsets([1, 2]) == [[], [1], [1, 2], [2]]


def test_solve_n_queens() -> None:
    assert len(solution.solve_n_queens(4)) == 2
    assert solution.solve_n_queens(4)[0] == [".Q..", "...Q", "Q...", "..Q."]
    assert solution.solve_n_queens(0) == []


@pytest.mark.parametrize("n", [1, 2, 3, 4, 5, 6])
def test_solve_n_queens_counts(n: int) -> None:
    expected = {1: 1, 2: 0, 3: 0, 4: 2, 5: 10, 6: 4}
    assert len(solution.solve_n_queens(n)) == expected[n]
