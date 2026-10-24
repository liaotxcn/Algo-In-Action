"""10 动态规划 —— Python 实现。"""

from __future__ import annotations

import bisect


def fib(n: int) -> int:
    """斐波那契数列（滚动变量），时间复杂度 O(n)。"""
    if n <= 0:
        return 0
    prev, current = 0, 1
    for _ in range(2, n + 1):
        prev, current = current, prev + current
    return current


def climb_stairs(n: int) -> int:
    """爬楼梯方案数。"""
    if n <= 2:
        return max(n, 0)
    a, b = 1, 2
    for _ in range(3, n + 1):
        a, b = b, a + b
    return b


def knapsack_01(weights: list[int], values: list[int], capacity: int) -> int:
    """0-1 背包最大价值，空间压缩到 O(C)。"""
    if capacity <= 0 or not weights:
        return 0
    dp = [0] * (capacity + 1)
    for weight, value in zip(weights, values):
        if weight > capacity:
            continue
        for current in range(capacity, weight - 1, -1):
            dp[current] = max(dp[current], dp[current - weight] + value)
    return dp[capacity]


def lis(nums: list[int]) -> int:
    """最长严格递增子序列，时间复杂度 O(n log n)。"""
    tails: list[int] = []
    for value in nums:
        index = bisect.bisect_left(tails, value)
        if index == len(tails):
            tails.append(value)
        else:
            tails[index] = value
    return len(tails)


def edit_distance(word1: str, word2: str) -> int:
    """编辑距离，空间复杂度 O(n)。"""
    if len(word1) < len(word2):
        word1, word2 = word2, word1

    previous = list(range(len(word2) + 1))
    current = [0] * (len(word2) + 1)
    for i in range(1, len(word1) + 1):
        current[0] = i
        for j in range(1, len(word2) + 1):
            if word1[i - 1] == word2[j - 1]:
                current[j] = previous[j - 1]
            else:
                current[j] = 1 + min(previous[j - 1], previous[j], current[j - 1])
        previous, current = current, previous
    return previous[len(word2)]


def coin_change(coins: list[int], amount: int) -> int:
    """最少硬币数，无法凑出返回 -1。"""
    unreachable = float("inf")
    dp = [0] + [unreachable] * amount
    for coin in coins:
        if coin <= 0:
            continue
        for value in range(coin, amount + 1):
            dp[value] = min(dp[value], dp[value - coin] + 1)
    return -1 if dp[amount] == unreachable else int(dp[amount])
