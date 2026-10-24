"""03 TopK 排行 —— Python 实现。"""

from __future__ import annotations

import heapq
from collections import Counter


def top_k(nums: list[int], k: int) -> list[int]:
    """返回最大的 k 个数，按降序排列。"""
    if k <= 0 or not nums:
        return []
    k = min(k, len(nums))

    top: list[int] = []
    for value in nums:
        if len(top) < k:
            heapq.heappush(top, value)
        elif value > top[0]:
            heapq.heapreplace(top, value)
    return sorted(top, reverse=True)


class _WordFreq:
    """按「更差者更小」定义顺序，便于用小顶堆维护 TopK。"""

    __slots__ = ("word", "count")

    def __init__(self, word: str, count: int) -> None:
        self.word = word
        self.count = count

    def __lt__(self, other: "_WordFreq") -> bool:
        if self.count != other.count:
            return self.count < other.count
        return self.word > other.word


def _is_better(candidate: _WordFreq, current: _WordFreq) -> bool:
    """判断 candidate 是否优于 current。"""
    if candidate.count != current.count:
        return candidate.count > current.count
    return candidate.word < current.word


def top_k_frequent(words: list[str], k: int) -> list[str]:
    """返回出现次数最多的 k 个单词；次数相同时字典序小的优先。"""
    if k <= 0 or not words:
        return []

    counts = Counter(words)
    top: list[_WordFreq] = []
    for word, count in counts.items():
        candidate = _WordFreq(word, count)
        if len(top) < k:
            heapq.heappush(top, candidate)
        elif _is_better(candidate, top[0]):
            heapq.heapreplace(top, candidate)

    top.sort(key=lambda item: (-item.count, item.word))
    return [item.word for item in top]
