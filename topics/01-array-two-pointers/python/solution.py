"""01 数组与双指针 —— Python 实现。"""

from __future__ import annotations

import json
from pathlib import Path


def rotate_array(nums: list[int], k: int) -> None:
    """原地向右轮转 k 步，时间复杂度 O(n)，空间复杂度 O(1)。"""
    n = len(nums)
    if n == 0:
        return
    k = ((k % n) + n) % n
    if k == 0:
        return
    nums.reverse()
    nums[:k] = reversed(nums[:k])
    nums[k:] = reversed(nums[k:])


def max_area(height: list[int]) -> int:
    """双指针求容器最大面积，时间复杂度 O(n)。"""
    ans = 0
    left, right = 0, len(height) - 1
    while left < right:
        ans = max(ans, (right - left) * min(height[left], height[right]))
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return ans


def three_sum(nums: list[int]) -> list[list[int]]:
    """返回所有和为 0 的不重复三元组，时间复杂度 O(n²)。"""
    n = len(nums)
    if n < 3:
        return []
    nums.sort()
    ans: list[list[int]] = []

    for i in range(n - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        if nums[i] + nums[i + 1] + nums[i + 2] > 0:
            break
        if nums[i] + nums[n - 2] + nums[n - 1] < 0:
            continue

        left, right = i + 1, n - 1
        while left < right:
            total = nums[i] + nums[left] + nums[right]
            if total > 0:
                right -= 1
            elif total < 0:
                left += 1
            else:
                ans.append([nums[i], nums[left], nums[right]])
                while left < right and nums[left] == nums[left + 1]:
                    left += 1
                while left < right and nums[right] == nums[right - 1]:
                    right -= 1
                left += 1
                right -= 1
    return ans


def move_zeroes(nums: list[int]) -> None:
    """原地移动零到末尾，保持非零元素顺序。"""
    write = 0
    for value in nums:
        if value != 0:
            nums[write] = value
            write += 1
    for index in range(write, len(nums)):
        nums[index] = 0


def trap(height: list[int]) -> int:
    """双指针计算接雨水总量，时间复杂度 O(n)。"""
    ans = 0
    left, right = 0, len(height) - 1
    pre_max = suf_max = 0
    while left < right:
        pre_max = max(pre_max, height[left])
        suf_max = max(suf_max, height[right])
        if pre_max < suf_max:
            ans += pre_max - height[left]
            left += 1
        else:
            ans += suf_max - height[right]
            right -= 1
    return ans


def merge_intervals(intervals: list[list[int]]) -> list[list[int]]:
    """合并重叠区间，不修改入参，时间复杂度 O(n log n)。"""
    if not intervals:
        return []
    ordered = sorted(intervals, key=lambda item: item[0])
    ans: list[list[int]] = []
    for start, end in ordered:
        if ans and start <= ans[-1][1]:
            ans[-1][1] = max(ans[-1][1], end)
        else:
            ans.append([start, end])
    return ans


def product_except_self(nums: list[int]) -> list[int]:
    """不使用除法计算除自身外的乘积，时间复杂度 O(n)。"""
    n = len(nums)
    if n == 0:
        return []
    suffix = [1] * n
    for i in range(n - 2, -1, -1):
        suffix[i] = suffix[i + 1] * nums[i + 1]
    prefix = 1
    for i, value in enumerate(nums):
        suffix[i] *= prefix
        prefix *= value
    return suffix


def max_sub_array(nums: list[int]) -> int:
    """最大子数组和（前缀和 + 最小前缀），时间复杂度 O(n)。"""
    if not nums:
        return 0
    ans = nums[0]
    prefix = min_prefix = 0
    for value in nums:
        prefix += value
        ans = max(ans, prefix - min_prefix)
        min_prefix = min(min_prefix, prefix)
    return ans


def spiral_order(matrix: list[list[int]]) -> list[int]:
    """顺时针螺旋遍历矩阵。"""
    if not matrix or not matrix[0]:
        return []
    top, bottom = 0, len(matrix) - 1
    left, right = 0, len(matrix[0]) - 1
    ans: list[int] = []

    while top <= bottom and left <= right:
        for col in range(left, right + 1):
            ans.append(matrix[top][col])
        top += 1
        for row in range(top, bottom + 1):
            ans.append(matrix[row][right])
        right -= 1
        if top <= bottom:
            for col in range(right, left - 1, -1):
                ans.append(matrix[bottom][col])
            bottom -= 1
        if left <= right:
            for row in range(bottom, top - 1, -1):
                ans.append(matrix[row][left])
            left += 1
    return ans


def set_zeroes(matrix: list[list[int]]) -> None:
    """原地矩阵置零，额外空间 O(1)。"""
    if not matrix or not matrix[0]:
        return
    rows, cols = len(matrix), len(matrix[0])
    first_row_has_zero = 0 in matrix[0]

    for row in range(1, rows):
        for col, value in enumerate(matrix[row]):
            if value == 0:
                matrix[row][0] = 0
                matrix[0][col] = 0

    for row in range(1, rows):
        for col in range(cols - 1, -1, -1):
            if matrix[row][0] == 0 or matrix[0][col] == 0:
                matrix[row][col] = 0

    if first_row_has_zero:
        matrix[0] = [0] * cols


def rotate_matrix(matrix: list[list[int]]) -> None:
    """原地顺时针旋转 n×n 矩阵 90 度。"""
    n = len(matrix)
    for i in range(n):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        matrix[i].reverse()


def search_matrix(matrix: list[list[int]], target: int) -> bool:
    """在行列均升序的矩阵中查找目标值，时间复杂度 O(m+n)。"""
    if not matrix or not matrix[0]:
        return False
    row, col = 0, len(matrix[0]) - 1
    while row < len(matrix) and col >= 0:
        value = matrix[row][col]
        if value == target:
            return True
        if value < target:
            row += 1
        else:
            col -= 1
    return False


def load_shared_cases() -> dict:
    """读取与 Go 测试共享的用例文件。"""
    path = Path(__file__).resolve().parent.parent / "cases" / "cases.json"
    return json.loads(path.read_text(encoding="utf-8"))
