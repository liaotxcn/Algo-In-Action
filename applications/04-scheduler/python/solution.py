"""04 任务调度与分布式 ID —— Python 实现。"""

from __future__ import annotations

import heapq
import threading
import time
from collections.abc import Callable
from datetime import UTC, datetime

NODE_BITS = 10
SEQUENCE_BITS = 12
MAX_NODE_ID = (1 << NODE_BITS) - 1
MAX_SEQUENCE = (1 << SEQUENCE_BITS) - 1
TIMESTAMP_SHIFT = NODE_BITS + SEQUENCE_BITS
NODE_SHIFT = SEQUENCE_BITS

# 自定义纪元：2020-01-01 UTC（毫秒）
EPOCH_MILLIS = 1577836800000


class SnowflakeError(ValueError):
    """Snowflake 生成器参数或状态错误。"""


class Snowflake:
    """64 位趋势递增 ID 生成器。"""

    def __init__(
        self,
        node_id: int,
        clock: Callable[[], float] = time.time,
    ) -> None:
        if node_id < 0 or node_id > MAX_NODE_ID:
            raise SnowflakeError("节点 ID 超出范围")
        self.node_id = node_id
        self.clock = clock
        self._last_millis = 0
        self._sequence = 0
        self._lock = threading.Lock()

    def _now_millis(self) -> int:
        """返回当前毫秒时间戳。"""
        return int(self.clock() * 1000)

    def next_id(self) -> int:
        """生成下一个 ID；时钟回拨或序列号耗尽时抛出 SnowflakeError。"""
        with self._lock:
            current = self._now_millis()
            if current < self._last_millis:
                raise SnowflakeError("检测到时钟回拨")
            if current == self._last_millis:
                if self._sequence >= MAX_SEQUENCE:
                    raise SnowflakeError("当前毫秒内序列号已用尽")
                self._sequence += 1
            else:
                self._sequence = 0
            self._last_millis = current

            return (
                ((current - EPOCH_MILLIS) << TIMESTAMP_SHIFT)
                | (self.node_id << NODE_SHIFT)
                | self._sequence
            )

    def parse(self, value: int) -> tuple[datetime, int, int]:
        """拆解 ID，返回生成时间、机器号与序列号。"""
        millis = (value >> TIMESTAMP_SHIFT) + EPOCH_MILLIS
        generated_at = datetime.fromtimestamp(millis / 1000, tz=UTC)
        node_id = (value >> NODE_SHIFT) & MAX_NODE_ID
        sequence = value & MAX_SEQUENCE
        return generated_at, node_id, sequence


class Scheduler:
    """轮询调度器。"""

    def __init__(self, workers: list[str]) -> None:
        self.workers = list(workers)
        self.next = 0
        self._lock = threading.Lock()

    def assign(self) -> str:
        """返回下一个执行节点；没有节点时抛出 LookupError。"""
        with self._lock:
            if not self.workers:
                raise LookupError("没有可用的执行节点")
            worker = self.workers[self.next]
            self.next = (self.next + 1) % len(self.workers)
            return worker


class PriorityScheduler:
    """优先级调度器：优先级高者先执行，同优先级保持提交顺序。"""

    def __init__(self) -> None:
        self._queue: list[tuple[int, int, str]] = []
        self._next_seq = 0
        self._lock = threading.Lock()

    def submit(self, name: str, priority: int) -> None:
        """提交任务。"""
        with self._lock:
            # 堆按 (优先级, 提交序号) 排序；优先级取负数实现大顶堆效果
            heapq.heappush(self._queue, (-priority, self._next_seq, name))
            self._next_seq += 1

    def pop(self) -> tuple[str, int] | None:
        """取出下一个任务；队列为空时返回 None。"""
        with self._lock:
            if not self._queue:
                return None
            negative_priority, _, name = heapq.heappop(self._queue)
            return name, -negative_priority
