"""08 流式统计 —— Python 实现。"""

from __future__ import annotations

import hashlib
import random
import threading
from collections import deque

DEFAULT_SKETCH_WIDTH = 256
DEFAULT_SKETCH_DEPTH = 4
HASH_SEED_STRIDE = 0x9E3779B97F4A7C15


def _hash_with_seed(data: bytes, seed: int) -> int:
    """返回带种子的 64 位哈希。"""
    digest = hashlib.blake2b(
        seed.to_bytes(8, "little") + data, digest_size=8
    ).digest()
    return int.from_bytes(digest, "little")


class CountMinSketch:
    """Count-Min Sketch：固定内存估计元素频次，只高估不低估。"""

    def __init__(self, width: int = DEFAULT_SKETCH_WIDTH, depth: int = DEFAULT_SKETCH_DEPTH) -> None:
        self.width = width if width > 0 else DEFAULT_SKETCH_WIDTH
        self.depth = depth if depth > 0 else DEFAULT_SKETCH_DEPTH
        self._counts = [[0] * self.width for _ in range(self.depth)]
        self._lock = threading.Lock()

    def _column(self, data: bytes, row: int) -> int:
        """返回某一行对应的桶下标。"""
        seed = (row * HASH_SEED_STRIDE + 1) & 0xFFFFFFFFFFFFFFFF
        return _hash_with_seed(data, seed) % self.width

    def add(self, data: bytes, delta: int = 1) -> None:
        """增加元素的计数。"""
        if delta <= 0:
            return
        with self._lock:
            for row in range(self.depth):
                self._counts[row][self._column(data, row)] += delta

    def estimate(self, data: bytes) -> int:
        """返回元素出现次数的估计值。"""
        with self._lock:
            return min(
                self._counts[row][self._column(data, row)]
                for row in range(self.depth)
            )


class MovingAverage:
    """最近 capacity 个数值的滑动平均。"""

    def __init__(self, capacity: int) -> None:
        self.capacity = max(1, capacity)
        self._values: deque[float] = deque()
        self._total = 0.0
        self._lock = threading.Lock()

    def add(self, value: float) -> None:
        """加入新数值，窗口已满时淘汰最旧值。"""
        with self._lock:
            if len(self._values) == self.capacity:
                self._total -= self._values.popleft()
            self._values.append(value)
            self._total += value

    @property
    def average(self) -> float:
        """当前窗口平均值；窗口为空时返回 0。"""
        with self._lock:
            if not self._values:
                return 0.0
            return self._total / len(self._values)

    def __len__(self) -> int:
        with self._lock:
            return len(self._values)


class ReservoirSampler:
    """蓄水池抽样：从数据流中等概率抽取固定数量样本。"""

    def __init__(self, capacity: int, seed: int) -> None:
        self.capacity = max(0, capacity)
        self._random = random.Random(seed)
        self._items: list[int] = []
        self._seen = 0
        self._lock = threading.Lock()

    def add(self, value: int) -> None:
        """向数据流中加入一个元素。"""
        with self._lock:
            self._seen += 1
            if len(self._items) < self.capacity:
                self._items.append(value)
                return
            if self.capacity == 0:
                return
            slot = self._random.randrange(self._seen)
            if slot < self.capacity:
                self._items[slot] = value

    @property
    def sample(self) -> list[int]:
        """当前样本的副本。"""
        with self._lock:
            return list(self._items)

    @property
    def seen(self) -> int:
        """已观察到的元素总数。"""
        with self._lock:
            return self._seen
