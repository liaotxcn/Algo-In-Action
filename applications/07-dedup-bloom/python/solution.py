"""07 去重过滤器 —— Python 实现。"""

from __future__ import annotations

import hashlib
import math
import threading

DEFAULT_FALSE_POSITIVE_RATE = 0.01
DEFAULT_BUCKET_SIZE = 4
DEFAULT_MAX_KICKS = 500
FINGERPRINT_MASK = 0xFFFF
FINGERPRINT_SEED = b"\x51\x51"
PRIMARY_SEED = b"\xa5\xa5"
ALTERNATE_SEED = b"\xbe\xef"


def _hash_bytes(data: bytes, seed: bytes = b"") -> int:
    """返回带种子的 64 位哈希。"""
    digest = hashlib.blake2b(seed + data, digest_size=8).digest()
    return int.from_bytes(digest, "little")


class BloomFilter:
    """线程安全的布隆过滤器。"""

    def __init__(
        self,
        expected_items: int,
        false_positive_rate: float = DEFAULT_FALSE_POSITIVE_RATE,
    ) -> None:
        expected_items = max(1, expected_items)
        if not 0 < false_positive_rate < 1:
            false_positive_rate = DEFAULT_FALSE_POSITIVE_RATE

        self.bit_count = math.ceil(
            -expected_items * math.log(false_positive_rate) / (math.log(2) ** 2)
        )
        self.hash_count = max(
            1, round(self.bit_count / expected_items * math.log(2))
        )
        self._bits = bytearray((self.bit_count + 7) // 8)
        self._item_count = 0
        self._lock = threading.RLock()

    def _positions(self, data: bytes) -> list[int]:
        """返回元素对应的 k 个位下标（双重哈希）。"""
        first = _hash_bytes(data)
        second = _hash_bytes(data, b"\x01") | 1
        return [
            (first + index * second) % self.bit_count
            for index in range(self.hash_count)
        ]

    def add(self, data: bytes) -> None:
        """把元素加入过滤器。"""
        with self._lock:
            for position in self._positions(data):
                self._bits[position // 8] |= 1 << (position % 8)
            self._item_count += 1

    def might_contain(self, data: bytes) -> bool:
        """判断元素是否可能存在；False 表示一定不存在。"""
        with self._lock:
            return all(
                self._bits[position // 8] & (1 << (position % 8))
                for position in self._positions(data)
            )

    @property
    def count(self) -> int:
        """已添加的元素个数。"""
        with self._lock:
            return self._item_count


class CuckooFilter:
    """支持删除的布谷鸟过滤器。"""

    def __init__(self, capacity: int) -> None:
        capacity = max(1, capacity)
        bucket_count = 1
        while bucket_count * DEFAULT_BUCKET_SIZE < capacity:
            bucket_count <<= 1
        bucket_count = max(2, bucket_count)

        self._buckets: list[list[int]] = [
            [0] * DEFAULT_BUCKET_SIZE for _ in range(bucket_count)
        ]
        self._mask = bucket_count - 1
        self._bucket_size = DEFAULT_BUCKET_SIZE
        self._max_kicks = DEFAULT_MAX_KICKS
        self._count = 0
        self._lock = threading.RLock()

    def _fingerprint(self, data: bytes) -> int:
        """返回非零 16 位指纹。"""
        fingerprint = _hash_bytes(data, FINGERPRINT_SEED) & FINGERPRINT_MASK
        return fingerprint or 1

    def _primary_index(self, data: bytes) -> int:
        """返回主桶下标。"""
        return _hash_bytes(data, PRIMARY_SEED) & self._mask

    def _alt_index(self, index: int, fingerprint: int) -> int:
        """返回备用桶下标。"""
        seed = fingerprint.to_bytes(2, "little")
        return (index ^ _hash_bytes(seed, ALTERNATE_SEED)) & self._mask

    def _insert_into(self, index: int, fingerprint: int) -> bool:
        """尝试把指纹放入指定桶的空槽。"""
        bucket = self._buckets[index]
        for slot, value in enumerate(bucket):
            if value == 0:
                bucket[slot] = fingerprint
                return True
        return False

    def add(self, data: bytes) -> bool:
        """插入元素；过滤器已满且踢出失败时返回 False。"""
        with self._lock:
            fingerprint = self._fingerprint(data)
            first = self._primary_index(data)
            second = self._alt_index(first, fingerprint)
            if self._insert_into(first, fingerprint) or self._insert_into(
                second, fingerprint
            ):
                self._count += 1
                return True

            index = first
            current = fingerprint
            for kick in range(self._max_kicks):
                slot = kick % self._bucket_size
                bucket = self._buckets[index]
                current, bucket[slot] = bucket[slot], current
                index = self._alt_index(index, current)
                if self._insert_into(index, current):
                    self._count += 1
                    return True
            return False

    def _contains_in(self, index: int, fingerprint: int) -> bool:
        """判断指定桶是否包含指纹。"""
        return fingerprint in self._buckets[index]

    def might_contain(self, data: bytes) -> bool:
        """判断元素是否存在；False 表示一定不存在。"""
        with self._lock:
            fingerprint = self._fingerprint(data)
            first = self._primary_index(data)
            second = self._alt_index(first, fingerprint)
            return self._contains_in(first, fingerprint) or self._contains_in(
                second, fingerprint
            )

    def delete(self, data: bytes) -> bool:
        """删除元素；不存在时返回 False。"""
        with self._lock:
            fingerprint = self._fingerprint(data)
            first = self._primary_index(data)
            second = self._alt_index(first, fingerprint)
            for index in (first, second):
                bucket = self._buckets[index]
                for slot, value in enumerate(bucket):
                    if value == fingerprint:
                        bucket[slot] = 0
                        self._count -= 1
                        return True
            return False

    @property
    def count(self) -> int:
        """已插入的元素个数。"""
        with self._lock:
            return self._count
