"""08 排序 —— Python 实现。

约定：以下函数返回升序新列表，不修改入参。
"""

from __future__ import annotations


def bubble_sort(nums: list[int]) -> list[int]:
    """冒泡排序，时间复杂度 O(n²)。"""
    result = nums.copy()
    for end in range(len(result) - 1, 0, -1):
        swapped = False
        for i in range(end):
            if result[i] > result[i + 1]:
                result[i], result[i + 1] = result[i + 1], result[i]
                swapped = True
        if not swapped:
            break
    return result


def selection_sort(nums: list[int]) -> list[int]:
    """选择排序，时间复杂度 O(n²)。"""
    result = nums.copy()
    for i in range(len(result) - 1):
        min_index = i
        for j in range(i + 1, len(result)):
            if result[j] < result[min_index]:
                min_index = j
        result[i], result[min_index] = result[min_index], result[i]
    return result


def insertion_sort(nums: list[int]) -> list[int]:
    """直接插入排序，最好时间复杂度 O(n)。"""
    result = nums.copy()
    for i in range(1, len(result)):
        key = result[i]
        j = i - 1
        while j >= 0 and result[j] > key:
            result[j + 1] = result[j]
            j -= 1
        result[j + 1] = key
    return result


def shell_sort(nums: list[int]) -> list[int]:
    """希尔排序，增量每次减半。"""
    result = nums.copy()
    gap = len(result) // 2
    while gap > 0:
        for i in range(gap, len(result)):
            key = result[i]
            j = i - gap
            while j >= 0 and result[j] > key:
                result[j + gap] = result[j]
                j -= gap
            result[j + gap] = key
        gap //= 2
    return result


def merge_sort(nums: list[int]) -> list[int]:
    """归并排序，时间复杂度 O(n log n)，空间复杂度 O(n)。"""
    if len(nums) < 2:
        return nums.copy()

    mid = len(nums) // 2
    left = merge_sort(nums[:mid])
    right = merge_sort(nums[mid:])
    merged: list[int] = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            merged.append(left[i])
            i += 1
        else:
            merged.append(right[j])
            j += 1
    merged.extend(left[i:])
    merged.extend(right[j:])
    return merged


def quick_sort(nums: list[int]) -> list[int]:
    """三路快速排序，平均时间复杂度 O(n log n)。"""
    if len(nums) < 2:
        return nums.copy()

    pivot = nums[len(nums) // 2]
    lower = [value for value in nums if value < pivot]
    equal = [value for value in nums if value == pivot]
    higher = [value for value in nums if value > pivot]
    return quick_sort(lower) + equal + quick_sort(higher)


def heap_sort(nums: list[int]) -> list[int]:
    """堆排序，时间复杂度 O(n log n)。"""
    import heapq

    heap = nums.copy()
    heapq.heapify(heap)
    return [heapq.heappop(heap) for _ in range(len(heap))]


def counting_sort(nums: list[int]) -> list[int]:
    """计数排序，适用于范围有限的整数。"""
    if len(nums) < 2:
        return nums.copy()
    low, high = min(nums), max(nums)
    counts = [0] * (high - low + 1)
    for value in nums:
        counts[value - low] += 1

    result: list[int] = []
    for offset, count in enumerate(counts):
        result.extend([offset + low] * count)
    return result


def bucket_sort(nums: list[int]) -> list[int]:
    """桶排序，数据分布均匀时接近 O(n)。"""
    if len(nums) < 2:
        return nums.copy()

    low, high = min(nums), max(nums)
    if low == high:
        return nums.copy()

    bucket_count = max(1, len(nums) // 5)
    buckets: list[list[int]] = [[] for _ in range(bucket_count)]
    span = (high - low + 1) / bucket_count
    for value in nums:
        index = min(int((value - low) / span), bucket_count - 1)
        buckets[index].append(value)

    result: list[int] = []
    for bucket in buckets:
        result.extend(insertion_sort(bucket))
    return result


def radix_sort(nums: list[int]) -> list[int]:
    """LSD 基数排序，支持负数。"""
    if len(nums) < 2:
        return nums.copy()

    low = min(nums)
    shifted = [value - low for value in nums]
    max_value = max(shifted)
    exp = 1
    while max_value // exp > 0:
        buckets: list[list[int]] = [[] for _ in range(10)]
        for value in shifted:
            buckets[(value // exp) % 10].append(value)
        shifted = [value for bucket in buckets for value in bucket]
        exp *= 10
    return [value + low for value in shifted]
