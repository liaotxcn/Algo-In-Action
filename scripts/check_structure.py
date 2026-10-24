#!/usr/bin/env python3
"""校验仓库目录结构是否符合仓库约定。

用法:
    python scripts/check_structure.py

校验通过返回 0，否则打印全部问题并返回 1。
"""

from __future__ import annotations

import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent

# 模块清单
TOPICS = [
    "01-array-two-pointers",
    "02-linked-list",
    "03-stack-queue",
    "04-hash-table",
    "05-tree",
    "06-heap",
    "07-graph",
    "08-sorting",
    "09-searching",
    "10-dp",
    "11-greedy",
    "12-backtracking",
    "13-string",
]

APPLICATIONS = [
    "01-rate-limiter",
    "02-lru-cache",
    "03-topk-ranking",
    "04-scheduler",
    "05-text-search",
    "06-graph-route",
    "07-dedup-bloom",
    "08-stream-stats",
    "09-cache-eviction",
]

REQUIRED_DOCS = [
    "docs/index.md",
    "docs/paths/application.md",
    "docs/matrix.md",
]

REQUIRED_SCRIPTS = [
    "scripts/check_structure.py",
    "scripts/gen_matrix.py",
    "scripts/run_all_tests.sh",
    ".github/workflows/test.yml",
]

# 重构前的旧目录/文件，必须彻底迁移
FORBIDDEN_PATHS = [
    "Hot100",
    "Golang",
    "Python",
    "Cockoo",
    "Redis_Bloom",
    "SQL_Bloom",
    "bloom.go",
]


def rel(path: Path) -> str:
    """返回相对仓库根目录的 POSIX 路径，便于输出。"""
    return path.relative_to(ROOT).as_posix()


def has_heading(path: Path) -> bool:
    """判断 Markdown 文件是否包含一级标题。"""
    try:
        for line in path.read_text(encoding="utf-8").splitlines():
            if line.startswith("# "):
                return True
    except OSError:
        return False
    return False


def check_readme(module_dir: Path, errors: list[str]) -> None:
    """校验模块 README.md 存在且包含标题。"""
    readme = module_dir / "README.md"
    if not readme.is_file():
        errors.append(f"{rel(module_dir)}: 缺少 README.md")
    elif not has_heading(readme):
        errors.append(f"{rel(readme)}: 缺少一级标题")


def check_topic(topic: str, errors: list[str]) -> None:
    """校验单个学习主题的目录与文件。"""
    topic_dir = ROOT / "topics" / topic
    if not topic_dir.is_dir():
        errors.append(f"topics/{topic}: 目录不存在")
        return

    check_readme(topic_dir, errors)

    go_dir = topic_dir / "go"
    py_dir = topic_dir / "python"
    if not go_dir.is_dir() and not py_dir.is_dir():
        errors.append(f"topics/{topic}: 至少需要一个 go/ 或 python/ 实现目录")

    if go_dir.is_dir():
        for name in ("solution.go", "solution_test.go"):
            if not (go_dir / name).is_file():
                errors.append(f"topics/{topic}/go: 缺少 {name}")

    if py_dir.is_dir():
        for name in ("solution.py", "test_solution.py"):
            if not (py_dir / name).is_file():
                errors.append(f"topics/{topic}/python: 缺少 {name}")

    cases_dir = topic_dir / "cases"
    if cases_dir.is_dir() and not (cases_dir / "cases.json").is_file():
        errors.append(f"topics/{topic}/cases: 存在 cases/ 时必须提供 cases.json")


def check_application(app: str, errors: list[str]) -> None:
    """校验单个应用案例的目录与文件。"""
    app_dir = ROOT / "applications" / app
    if not app_dir.is_dir():
        errors.append(f"applications/{app}: 目录不存在")
        return

    check_readme(app_dir, errors)

    go_dir = app_dir / "go"
    py_dir = app_dir / "python"
    if not go_dir.is_dir() and not py_dir.is_dir():
        errors.append(f"applications/{app}: 至少需要一个 go/ 或 python/ 实现目录")

    if go_dir.is_dir():
        go_files = [p for p in go_dir.glob("*.go") if not p.name.endswith("_test.go")]
        test_files = list(go_dir.glob("*_test.go"))
        if not go_files:
            errors.append(f"applications/{app}/go: 缺少 Go 实现文件")
        if not test_files:
            errors.append(f"applications/{app}/go: 缺少 *_test.go")

    if py_dir.is_dir():
        py_files = [
            p
            for p in py_dir.glob("*.py")
            if not p.name.startswith("test_") and p.name != "conftest.py"
        ]
        test_files = list(py_dir.glob("test_*.py"))
        if not py_files:
            errors.append(f"applications/{app}/python: 缺少 Python 实现文件")
        if not test_files:
            errors.append(f"applications/{app}/python: 缺少 test_*.py")


def check_exact_dirs(parent: str, expected: list[str], errors: list[str]) -> None:
    """校验模块目录集合与规格完全一致，避免遗留未迁移目录。"""
    base = ROOT / parent
    if not base.is_dir():
        errors.append(f"{parent}/: 目录不存在")
        return

    actual = sorted(p.name for p in base.iterdir() if p.is_dir())
    missing = sorted(set(expected) - set(actual))
    extra = sorted(set(actual) - set(expected))
    for name in missing:
        errors.append(f"{parent}/{name}: 规格要求的目录不存在")
    for name in extra:
        errors.append(f"{parent}/{name}: 规格之外的目录，请迁移或删除")


def main() -> int:
    errors: list[str] = []

    check_exact_dirs("topics", TOPICS, errors)
    check_exact_dirs("applications", APPLICATIONS, errors)

    for topic in TOPICS:
        if (ROOT / "topics" / topic).is_dir():
            check_topic(topic, errors)

    for app in APPLICATIONS:
        if (ROOT / "applications" / app).is_dir():
            check_application(app, errors)

    for doc in REQUIRED_DOCS:
        if not (ROOT / doc).is_file():
            errors.append(f"{doc}: 缺失")

    for script in REQUIRED_SCRIPTS:
        if not (ROOT / script).is_file():
            errors.append(f"{script}: 缺失")

    for stale in FORBIDDEN_PATHS:
        if (ROOT / stale).exists():
            errors.append(f"{stale}: 旧结构仍存在，请迁移到 topics/ 或 applications/")

    if errors:
        print("目录结构校验失败：", file=sys.stderr)
        for err in errors:
            print(f"  - {err}", file=sys.stderr)
        return 1

    print(
        f"目录结构校验通过：{len(TOPICS)} 个学习主题，"
        f"{len(APPLICATIONS)} 个应用案例。"
    )
    return 0


if __name__ == "__main__":
    sys.exit(main())
