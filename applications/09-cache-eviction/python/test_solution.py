"""09 缓存淘汰策略的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_fifo_eviction() -> None:
    cache = solution.FIFOCache(2)
    cache.put(1, "one")
    cache.put(2, "two")
    assert cache.get(1) == "one"

    cache.put(3, "three")  # FIFO 只看写入顺序，淘汰 1
    assert cache.get(1) is None
    assert cache.get(2) == "two"
    assert cache.get(3) == "three"
    assert cache.evictions == 1


def test_fifo_update_does_not_reorder() -> None:
    cache = solution.FIFOCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    cache.put(1, 10)
    cache.put(3, 3)  # 仍淘汰 1

    assert cache.get(1) is None
    assert cache.get(2) == 2


def test_lfu_eviction() -> None:
    cache = solution.LFUCache(2)
    cache.put(1, "one")
    cache.put(2, "two")
    cache.get(1)  # 1 的频次变为 2
    cache.put(3, "three")  # 淘汰频次最低的 2

    assert cache.get(2) is None
    assert cache.get(1) == "one"
    assert cache.get(3) == "three"
    assert cache.evictions == 1


def test_lfu_tie_break_lru() -> None:
    cache = solution.LFUCache(2)
    cache.put("a", 1)
    cache.put("b", 2)
    cache.get("a")
    cache.get("b")
    cache.put("c", 3)  # 频次相同，淘汰更久未使用的 a

    assert cache.get("a") is None
    assert cache.get("b") == 2


def test_zero_capacity() -> None:
    fifo = solution.FIFOCache(0)
    fifo.put(1, 1)
    assert len(fifo) == 0

    lfu = solution.LFUCache(0)
    lfu.put(1, 1)
    assert len(lfu) == 0
