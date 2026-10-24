"""01 限流器的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


class FakeClock:
    """可手动推进的测试时钟。"""

    def __init__(self, current: float = 0.0) -> None:
        self.current = current

    def __call__(self) -> float:
        return self.current

    def advance(self, duration: float) -> None:
        self.current += duration


def test_token_bucket() -> None:
    clock = FakeClock()
    bucket = solution.TokenBucket(2, 1, clock)

    assert bucket.allow()
    assert bucket.allow()
    assert not bucket.allow()

    clock.advance(0.5)
    assert not bucket.allow()

    clock.advance(0.5)
    assert bucket.allow()


def test_token_bucket_refill_capped() -> None:
    clock = FakeClock()
    bucket = solution.TokenBucket(2, 10, clock)
    assert bucket.allow()
    assert bucket.allow()

    clock.advance(3600)
    assert bucket.tokens == 2


def test_sliding_window() -> None:
    clock = FakeClock()
    limiter = solution.SlidingWindow(2, 1.0, clock)

    assert limiter.allow()
    assert limiter.allow()
    assert not limiter.allow()

    clock.advance(0.4)
    assert not limiter.allow()

    clock.advance(0.7)
    assert limiter.allow()
    assert limiter.count == 1


def test_sliding_window_zero_limit() -> None:
    clock = FakeClock()
    limiter = solution.SlidingWindow(0, 1.0, clock)
    assert not limiter.allow()
