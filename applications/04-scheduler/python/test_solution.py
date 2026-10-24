"""04 任务调度的 pytest 测试。"""

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)


class StepClock:
    """可手动推进或回拨的测试时钟。"""

    def __init__(self, current: float = 1_700_000_000.0) -> None:
        self.current = current

    def __call__(self) -> float:
        return self.current

    def advance(self, duration: float) -> None:
        self.current += duration


def test_snowflake_ids() -> None:
    clock = StepClock()
    generator = solution.Snowflake(7, clock)

    first = generator.next_id()
    second = generator.next_id()
    assert second > first

    generated_at, node_id, sequence = generator.parse(second)
    assert node_id == 7
    assert sequence == 1
    assert generated_at.timestamp() == pytest.approx(clock.current)

    clock.advance(0.001)
    third = generator.next_id()
    assert generator.parse(third)[2] == 0


def test_snowflake_validation() -> None:
    with pytest.raises(solution.SnowflakeError):
        solution.Snowflake(-1)

    clock = StepClock()
    generator = solution.Snowflake(1, clock)
    generator.next_id()
    clock.advance(-1)
    with pytest.raises(solution.SnowflakeError):
        generator.next_id()


def test_round_robin_scheduler() -> None:
    scheduler = solution.Scheduler(["a", "b", "c"])
    assert [scheduler.assign() for _ in range(4)] == ["a", "b", "c", "a"]

    with pytest.raises(LookupError):
        solution.Scheduler([]).assign()


def test_priority_scheduler() -> None:
    scheduler = solution.PriorityScheduler()
    scheduler.submit("low", 1)
    scheduler.submit("high-1", 10)
    scheduler.submit("mid", 5)
    scheduler.submit("high-2", 10)

    assert scheduler.pop() == ("high-1", 10)
    assert scheduler.pop() == ("high-2", 10)
    assert scheduler.pop() == ("mid", 5)
    assert scheduler.pop() == ("low", 1)
    assert scheduler.pop() is None
