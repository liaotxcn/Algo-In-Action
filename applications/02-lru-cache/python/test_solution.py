"""02 LRU 缓存的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_get_put() -> None:
    cache = solution.LRUCache(2)
    cache.put(1, "one")
    cache.put(2, "two")

    assert cache.get(1) == "one"

    cache.put(3, "three")  # 淘汰最久未使用的 2
    assert cache.get(2) is None
    assert cache.get(3) == "three"
    assert cache.stats == (2, 1)


def test_update_refreshes_order() -> None:
    cache = solution.LRUCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    cache.put(1, 10)
    cache.put(3, 3)  # 应淘汰 2

    assert cache.get(1) == 10
    assert cache.get(2) is None
    assert len(cache) == 2


def test_zero_capacity() -> None:
    cache = solution.LRUCache(0)
    cache.put(1, 1)
    assert cache.get(1) is None
