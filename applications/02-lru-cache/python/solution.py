"""02 LRU 缓存 —— Python 实现。"""

from __future__ import annotations

import threading
from collections import OrderedDict
from typing import Generic, TypeVar

K = TypeVar("K")
V = TypeVar("V")


class LRUCache(Generic[K, V]):
    """基于 OrderedDict 的线程安全 LRU 缓存。"""

    def __init__(self, capacity: int) -> None:
        self.capacity = max(0, capacity)
        self._items: OrderedDict[K, V] = OrderedDict()
        self._lock = threading.Lock()
        self._hits = 0
        self._misses = 0

    def get(self, key: K, default: V | None = None) -> V | None:
        """读取键值；命中时把键移到最近使用位置。"""
        with self._lock:
            if key not in self._items:
                self._misses += 1
                return default
            self._hits += 1
            self._items.move_to_end(key)
            return self._items[key]

    def put(self, key: K, value: V) -> None:
        """写入键值；容量不足时淘汰最久未使用的键。"""
        if self.capacity == 0:
            return
        with self._lock:
            if key in self._items:
                self._items[key] = value
                self._items.move_to_end(key)
                return
            if len(self._items) >= self.capacity:
                self._items.popitem(last=False)
            self._items[key] = value

    def __len__(self) -> int:
        with self._lock:
            return len(self._items)

    @property
    def stats(self) -> tuple[int, int]:
        """返回 (命中次数, 未命中次数)。"""
        with self._lock:
            return self._hits, self._misses
