"""01 限流器 —— Python 实现。"""

from __future__ import annotations

import threading
import time
from collections.abc import Callable
from dataclasses import dataclass, field

Clock = Callable[[], float]

TOKENS_PER_REQUEST = 1.0
DEFAULT_CAPACITY = 1


@dataclass
class TokenBucket:
    """令牌桶限流器。"""

    capacity: int
    refill_per_second: float
    clock: Clock = time.monotonic
    _tokens: float = field(init=False)
    _last_refill: float = field(init=False)
    _lock: threading.Lock = field(default_factory=threading.Lock, init=False, repr=False)

    def __post_init__(self) -> None:
        if self.capacity < DEFAULT_CAPACITY:
            self.capacity = DEFAULT_CAPACITY
        self.refill_per_second = max(0.0, self.refill_per_second)
        self._tokens = float(self.capacity)
        self._last_refill = self.clock()

    def allow(self) -> bool:
        """尝试获取一个令牌，成功返回 True。"""
        with self._lock:
            self._refill_locked(self.clock())
            if self._tokens < TOKENS_PER_REQUEST:
                return False
            self._tokens -= TOKENS_PER_REQUEST
            return True

    @property
    def tokens(self) -> float:
        """当前可用令牌数。"""
        with self._lock:
            self._refill_locked(self.clock())
            return self._tokens

    def _refill_locked(self, current: float) -> None:
        """按经过时间补充令牌，调用方必须持有锁。"""
        elapsed = current - self._last_refill
        if elapsed > 0:
            self._tokens = min(
                self.capacity,
                self._tokens + elapsed * self.refill_per_second,
            )
            self._last_refill = current


class SlidingWindow:
    """滑动窗口限流器。"""

    def __init__(
        self,
        limit: int,
        window: float,
        clock: Clock = time.monotonic,
    ) -> None:
        self.limit = max(0, limit)
        self.window = max(0.0, window)
        self.clock = clock
        self._events: list[float] = []
        self._lock = threading.Lock()

    def allow(self) -> bool:
        """记录一次请求；窗口内请求数达到上限时返回 False。"""
        with self._lock:
            current = self.clock()
            cutoff = current - self.window
            self._events = [event for event in self._events if event > cutoff]
            if len(self._events) >= self.limit:
                return False
            self._events.append(current)
            return True

    @property
    def count(self) -> int:
        """当前窗口内的请求数。"""
        with self._lock:
            cutoff = self.clock() - self.window
            return sum(1 for event in self._events if event > cutoff)
