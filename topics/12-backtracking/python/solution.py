"""12 回溯 —— Python 实现。"""

from __future__ import annotations


def permutations(nums: list[int]) -> list[list[int]]:
    """返回不重复元素的全排列。"""
    result: list[list[int]] = []
    path: list[int] = []
    used = [False] * len(nums)

    def backtrack() -> None:
        if len(path) == len(nums):
            result.append(path.copy())
            return
        for index, value in enumerate(nums):
            if used[index]:
                continue
            used[index] = True
            path.append(value)
            backtrack()
            path.pop()
            used[index] = False

    backtrack()
    return result


def permutations_unique(nums: list[int]) -> list[list[int]]:
    """返回含重复元素的全排列（去重）。"""
    values = sorted(nums)
    result: list[list[int]] = []
    path: list[int] = []
    used = [False] * len(values)

    def backtrack() -> None:
        if len(path) == len(values):
            result.append(path.copy())
            return
        for index, value in enumerate(values):
            if used[index]:
                continue
            if index > 0 and values[index] == values[index - 1] and not used[index - 1]:
                continue
            used[index] = True
            path.append(value)
            backtrack()
            path.pop()
            used[index] = False

    backtrack()
    return result


def combine(n: int, k: int) -> list[list[int]]:
    """返回 1 到 n 中所有长度为 k 的组合。"""
    if k <= 0 or k > n:
        return []
    result: list[list[int]] = []
    path: list[int] = []

    def backtrack(start: int) -> None:
        if len(path) == k:
            result.append(path.copy())
            return
        for value in range(start, n - (k - len(path)) + 2):
            path.append(value)
            backtrack(value + 1)
            path.pop()

    backtrack(1)
    return result


def subsets(nums: list[int]) -> list[list[int]]:
    """返回全部子集（幂集）。"""
    result: list[list[int]] = []
    path: list[int] = []

    def backtrack(start: int) -> None:
        result.append(path.copy())
        for index in range(start, len(nums)):
            path.append(nums[index])
            backtrack(index + 1)
            path.pop()

    backtrack(0)
    return result


def solve_n_queens(n: int) -> list[list[str]]:
    """返回 n 皇后问题的所有解。"""
    if n <= 0:
        return []

    result: list[list[str]] = []
    cols = [False] * n
    diag = [False] * (2 * n - 1)
    anti_diag = [False] * (2 * n - 1)
    queens = [-1] * n

    def backtrack(row: int) -> None:
        if row == n:
            result.append(
                [
                    "." * col + "Q" + "." * (n - col - 1)
                    for col in queens
                ]
            )
            return
        for col in range(n):
            if cols[col] or diag[row + col] or anti_diag[row - col + n - 1]:
                continue
            cols[col] = diag[row + col] = anti_diag[row - col + n - 1] = True
            queens[row] = col
            backtrack(row + 1)
            queens[row] = -1
            cols[col] = diag[row + col] = anti_diag[row - col + n - 1] = False

    backtrack(0)
    return result
