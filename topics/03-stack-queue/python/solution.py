"""03 栈与队列 —— Python 实现。"""

from __future__ import annotations

from collections import deque


def is_valid(s: str) -> bool:
    """判断括号字符串是否有效。"""
    pairs = {")": "(", "]": "[", "}": "{"}
    stack: list[str] = []
    for ch in s:
        if ch in "([{":
            stack.append(ch)
        elif ch in pairs:
            if not stack or stack.pop() != pairs[ch]:
                return False
    return not stack


class MinStack:
    """支持 O(1) 获取最小值的栈。"""

    def __init__(self) -> None:
        self._data: list[int] = []
        self._mins: list[int] = []

    def push(self, value: int) -> None:
        self._data.append(value)
        if not self._mins or value <= self._mins[-1]:
            self._mins.append(value)

    def pop(self) -> int:
        if not self._data:
            return 0
        value = self._data.pop()
        if value == self._mins[-1]:
            self._mins.pop()
        return value

    def top(self) -> int:
        return self._data[-1] if self._data else 0

    def get_min(self) -> int:
        return self._mins[-1] if self._mins else 0


class MyQueue:
    """使用两个栈实现的 FIFO 队列。"""

    def __init__(self) -> None:
        self._in: list[int] = []
        self._out: list[int] = []

    def push(self, value: int) -> None:
        self._in.append(value)

    def pop(self) -> int:
        self._move()
        return self._out.pop() if self._out else 0

    def peek(self) -> int:
        self._move()
        return self._out[-1] if self._out else 0

    def empty(self) -> bool:
        return not self._in and not self._out

    def _move(self) -> None:
        if not self._out:
            while self._in:
                self._out.append(self._in.pop())


def max_sliding_window(nums: list[int], k: int) -> list[int]:
    """单调递减队列求滑动窗口最大值，时间复杂度 O(n)。"""
    if not nums or k <= 0:
        return []
    k = min(k, len(nums))

    candidates: deque[int] = deque()
    ans: list[int] = []
    for index, value in enumerate(nums):
        if candidates and candidates[0] <= index - k:
            candidates.popleft()
        while candidates and nums[candidates[-1]] <= value:
            candidates.pop()
        candidates.append(index)
        if index >= k - 1:
            ans.append(nums[candidates[0]])
    return ans
