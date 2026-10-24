"""09 搜索 —— Python 实现。"""

from __future__ import annotations


def binary_search(nums: list[int], target: int) -> int:
    """二分查找，返回任意匹配下标；未找到返回 -1。"""
    left, right = 0, len(nums) - 1
    while left <= right:
        mid = left + (right - left) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    return -1


def lower_bound(nums: list[int], target: int) -> int:
    """返回第一个 >= target 的下标。"""
    left, right = 0, len(nums)
    while left < right:
        mid = left + (right - left) // 2
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid
    return left


def upper_bound(nums: list[int], target: int) -> int:
    """返回第一个 > target 的下标。"""
    left, right = 0, len(nums)
    while left < right:
        mid = left + (right - left) // 2
        if nums[mid] <= target:
            left = mid + 1
        else:
            right = mid
    return left


def search_range(nums: list[int], target: int) -> tuple[int, int]:
    """返回 target 的起止下标，不存在返回 (-1, -1)。"""
    start = lower_bound(nums, target)
    if start == len(nums) or nums[start] != target:
        return (-1, -1)
    return (start, upper_bound(nums, target) - 1)


def interpolation_search(nums: list[int], target: int) -> int:
    """插值查找，适用于分布均匀的有序数组。"""
    left, right = 0, len(nums) - 1
    while left <= right and nums[left] <= target <= nums[right]:
        if nums[left] == nums[right]:
            return left if nums[left] == target else -1

        pos = left + (target - nums[left]) * (right - left) // (
            nums[right] - nums[left]
        )
        if nums[pos] == target:
            return pos
        if nums[pos] < target:
            left = pos + 1
        else:
            right = pos - 1
    return -1
