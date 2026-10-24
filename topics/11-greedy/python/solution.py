"""11 贪心 —— Python 实现。"""

from __future__ import annotations


def activity_selection(start: list[int], end: list[int]) -> list[int]:
    """按结束时间贪心选择最多不重叠活动。"""
    if not start or len(start) != len(end):
        return []
    order = sorted(range(len(start)), key=lambda index: end[index])
    selected: list[int] = []
    last_end = float("-inf")
    for index in order:
        if start[index] >= last_end:
            selected.append(index)
            last_end = end[index]
    return selected


def fractional_knapsack(weights: list[int], values: list[int], capacity: float) -> float:
    """分数背包最大价值。"""
    if capacity <= 0 or not weights or len(weights) != len(values):
        return 0.0

    items = [
        (values[i] / weight if weight > 0 else 0.0, weight)
        for i, weight in enumerate(weights)
    ]
    items.sort(key=lambda item: item[0], reverse=True)

    total = 0.0
    remaining = capacity
    for ratio, weight in items:
        if remaining <= 0:
            break
        take = min(weight, remaining)
        total += take * ratio
        remaining -= take
    return total


def jump_game(nums: list[int]) -> bool:
    """跳跃游戏可达性判断。"""
    reach = 0
    for index, jump in enumerate(nums):
        if index > reach:
            return False
        reach = max(reach, index + jump)
    return True


def min_jumps(nums: list[int]) -> int:
    """跳跃游戏最少步数；不可达返回 -1。"""
    if len(nums) <= 1:
        return 0
    jumps = current_end = farthest = 0
    for index in range(len(nums) - 1):
        farthest = max(farthest, index + nums[index])
        if index == current_end:
            jumps += 1
            current_end = farthest
            if current_end >= len(nums) - 1:
                return jumps
    return -1


def find_content_children(greed: list[int], cookies: list[int]) -> int:
    """贪心分配饼干，返回最多满足的孩子数量。"""
    greed = sorted(greed)
    cookies = sorted(cookies)
    child = cookie = 0
    while child < len(greed) and cookie < len(cookies):
        if cookies[cookie] >= greed[child]:
            child += 1
        cookie += 1
    return child
