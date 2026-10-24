"""07 去重过滤器的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_bloom_filter_no_false_negative() -> None:
    bloom = solution.BloomFilter(100, 0.01)
    items = [f"item-{index}" for index in range(100)]
    for item in items:
        bloom.add(item.encode())

    assert all(bloom.might_contain(item.encode()) for item in items)
    assert bloom.count == len(items)


def test_bloom_filter_false_positive_rate() -> None:
    bloom = solution.BloomFilter(100, 0.01)
    for index in range(100):
        bloom.add(f"item-{index}".encode())

    queries = 1000
    false_positives = sum(
        bloom.might_contain(f"absent-{index}".encode()) for index in range(queries)
    )
    # 理论误判率约 1%，放宽到 10% 以避免测试不稳定
    assert false_positives <= queries // 10


def test_cuckoo_filter() -> None:
    cuckoo = solution.CuckooFilter(256)
    items = [f"item-{index}" for index in range(100)]
    for item in items:
        assert cuckoo.add(item.encode())

    assert all(cuckoo.might_contain(item.encode()) for item in items)
    assert cuckoo.count == len(items)


def test_cuckoo_filter_delete() -> None:
    cuckoo = solution.CuckooFilter(64)
    cuckoo.add(b"alpha")
    cuckoo.add(b"beta")

    assert cuckoo.delete(b"alpha")
    assert not cuckoo.might_contain(b"alpha")
    assert cuckoo.might_contain(b"beta")
    assert not cuckoo.delete(b"alpha")
    assert cuckoo.count == 1
