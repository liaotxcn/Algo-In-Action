"""06 堆 —— Python 实现。"""

from __future__ import annotations

import heapq


class MaxHeap:
    """大顶堆（内部用负数存入 heapq 实现）。"""

    def __init__(self) -> None:
        self._data: list[int] = []

    def __len__(self) -> int:
        return len(self._data)

    def push(self, value: int) -> None:
        heapq.heappush(self._data, -value)

    def pop(self) -> int | None:
        return -heapq.heappop(self._data) if self._data else None

    def peek(self) -> int | None:
        return -self._data[0] if self._data else None


def heap_sort(nums: list[int]) -> list[int]:
    """堆排序，返回升序新列表。"""
    heap = nums.copy()
    heapq.heapify(heap)
    return [heapq.heappop(heap) for _ in range(len(heap))]


def top_k(nums: list[int], k: int) -> list[int]:
    """返回最大的 k 个元素（从大到小）。"""
    if k <= 0 or not nums:
        return []
    return sorted(heapq.nlargest(min(k, len(nums)), nums), reverse=True)
