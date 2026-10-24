"""09 缓存淘汰策略 —— Python 实现。"""

from __future__ import annotations

import threading
from collections import OrderedDict
from typing import Generic, TypeVar

K = TypeVar("K")
V = TypeVar("V")


class FIFOCache(Generic[K, V]):
    """按写入顺序淘汰数据的线程安全缓存。"""

    def __init__(self, capacity: int) -> None:
        self.capacity = max(0, capacity)
        self._items: OrderedDict[K, V] = OrderedDict()
        self._evictions = 0
        self._lock = threading.Lock()

    def get(self, key: K, default: V | None = None) -> V | None:
        """读取键值；不改变淘汰顺序。"""
        with self._lock:
            return self._items.get(key, default)

    def put(self, key: K, value: V) -> None:
        """写入键值；容量不足时淘汰最早写入的键。"""
        if self.capacity == 0:
            return
        with self._lock:
            if key in self._items:
                self._items[key] = value
                return
            if len(self._items) >= self.capacity:
                self._items.popitem(last=False)
                self._evictions += 1
            self._items[key] = value

    def __len__(self) -> int:
        with self._lock:
            return len(self._items)

    @property
    def evictions(self) -> int:
        """累计淘汰次数。"""
        with self._lock:
            return self._evictions


class LFUCache(Generic[K, V]):
    """按访问频次淘汰数据的线程安全缓存，同频次淘汰最久未使用。"""

    def __init__(self, capacity: int) -> None:
        self.capacity = max(0, capacity)
        self._values: dict[K, V] = {}
        self._freqs: dict[K, int] = {}
        self._buckets: dict[int, OrderedDict[K, None]] = {}
        self._min_freq = 0
        self._evictions = 0
        self._lock = threading.Lock()

    def _touch(self, key: K) -> None:
        """把键移动到频次 +1 的桶尾部，调用方必须持有锁。"""
        freq = self._freqs[key]
        bucket = self._buckets[freq]
        del bucket[key]
        if not bucket:
            del self._buckets[freq]
            if self._min_freq == freq:
                self._min_freq += 1
        freq += 1
        self._freqs[key] = freq
        self._buckets.setdefault(freq, OrderedDict())[key] = None

    def get(self, key: K, default: V | None = None) -> V | None:
        """读取键值；命中时提升访问频次。"""
        with self._lock:
            if key not in self._values:
                return default
            self._touch(key)
            return self._values[key]

    def put(self, key: K, value: V) -> None:
        """写入键值；容量不足时淘汰频次最低的键。"""
        if self.capacity == 0:
            return
        with self._lock:
            if key in self._values:
                self._values[key] = value
                self._touch(key)
                return
            if len(self._values) >= self.capacity:
                bucket = self._buckets[self._min_freq]
                victim, _ = bucket.popitem(last=False)
                if not bucket:
                    del self._buckets[self._min_freq]
                del self._values[victim]
                del self._freqs[victim]
                self._evictions += 1

            self._values[key] = value
            self._freqs[key] = 1
            self._buckets.setdefault(1, OrderedDict())[key] = None
            self._min_freq = 1

    def __len__(self) -> int:
        with self._lock:
            return len(self._values)

    @property
    def evictions(self) -> int:
        """累计淘汰次数。"""
        with self._lock:
            return self._evictions
