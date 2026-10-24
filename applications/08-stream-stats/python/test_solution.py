"""08 流式统计的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_count_min_sketch() -> None:
    sketch = solution.CountMinSketch(256, 4)
    inputs = {"apple": 10, "banana": 3, "cherry": 7}
    for item, count in inputs.items():
        sketch.add(item.encode(), count)

    for item, count in inputs.items():
        assert sketch.estimate(item.encode()) >= count
    assert sketch.estimate(b"missing") == 0


def test_count_min_sketch_monotonic() -> None:
    sketch = solution.CountMinSketch(64, 3)
    sketch.add(b"key", 2)
    before = sketch.estimate(b"key")
    sketch.add(b"key", 5)
    assert sketch.estimate(b"key") >= before + 5


def test_moving_average() -> None:
    average = solution.MovingAverage(3)
    average.add(1)
    average.add(2)
    average.add(3)
    assert average.average == 2

    average.add(6)
    assert average.average == 11.0 / 3.0
    assert len(average) == 3


def test_moving_average_empty() -> None:
    assert solution.MovingAverage(2).average == 0


def test_reservoir_sampler_deterministic() -> None:
    first = solution.ReservoirSampler(5, 42)
    second = solution.ReservoirSampler(5, 42)
    for value in range(100):
        first.add(value)
        second.add(value)

    assert first.sample == second.sample
    assert first.seen == 100


def test_reservoir_sampler_capacity() -> None:
    sampler = solution.ReservoirSampler(3, 1)
    sampler.add(10)
    sampler.add(20)
    sampler.add(30)
    assert sampler.sample == [10, 20, 30]

    for value in range(40, 100):
        sampler.add(value)
    assert len(sampler.sample) == 3
    assert all(10 <= value < 100 for value in sampler.sample)
