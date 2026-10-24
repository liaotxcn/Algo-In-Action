"""08 排序的 pytest 测试，使用 shared/testdata 中的共享用例。"""

import json
from pathlib import Path

import pytest
from solution_loader import load_solution

solution = load_solution(__file__)

CASES_PATH = Path(__file__).resolve().parents[3] / "shared" / "testdata" / "sort_cases.json"
CASES = json.loads(CASES_PATH.read_text(encoding="utf-8"))["cases"]

SORTS = [
    solution.bubble_sort,
    solution.selection_sort,
    solution.insertion_sort,
    solution.shell_sort,
    solution.merge_sort,
    solution.quick_sort,
    solution.heap_sort,
    solution.counting_sort,
    solution.bucket_sort,
    solution.radix_sort,
]


@pytest.mark.parametrize("sort_fn", SORTS, ids=lambda fn: fn.__name__)
@pytest.mark.parametrize("case", CASES, ids=lambda case: case["name"])
def test_sort(sort_fn, case) -> None:
    original = list(case["input"])
    assert sort_fn(case["input"]) == case["expected"]
    assert case["input"] == original  # 不修改入参
