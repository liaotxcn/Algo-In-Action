#!/usr/bin/env python3
"""生成 docs/matrix.md 覆盖矩阵。

用法:
    python scripts/gen_matrix.py           # 写入 docs/matrix.md
    python scripts/gen_matrix.py --check   # 仅校验矩阵是否为最新
"""

from __future__ import annotations

import os
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
MATRIX_PATH = ROOT / "docs" / "matrix.md"

# 各语言的实现与测试文件探测规则
LANG_RULES = {
    "Go": {
        "impl": lambda d: [d / "solution.go"] if (d / "solution.go").is_file() else sorted(d.glob("*.go")),
        "test": lambda d: sorted(d.glob("*_test.go")),
    },
    "Python": {
        "impl": lambda d: [d / "solution.py"] if (d / "solution.py").is_file() else sorted(
            p for p in d.glob("*.py") if not p.name.startswith("test_") and p.name != "conftest.py"
        ),
        "test": lambda d: sorted(d.glob("test_*.py")),
    },
}

LANG_DIRS = {"Go": "go", "Python": "python"}


def rel_link(path: Path) -> str:
    """生成从 docs/ 指向目标文件的相对链接。"""
    return Path(os.path.relpath(path, ROOT / "docs")).as_posix()


def module_title(module_dir: Path) -> str:
    """读取 README 的一级标题，缺失时回退为目录名。"""
    readme = module_dir / "README.md"
    if readme.is_file():
        for line in readme.read_text(encoding="utf-8").splitlines():
            if line.startswith("# "):
                return line[2:].strip()
    return module_dir.name


def language_cells(module_dir: Path, lang: str) -> tuple[str, str]:
    """返回某语言的（实现单元格, 测试单元格）。"""
    lang_dir = module_dir / LANG_DIRS[lang]
    if not lang_dir.is_dir():
        return "—", "—"

    impl_files = LANG_RULES[lang]["impl"](lang_dir)
    test_files = LANG_RULES[lang]["test"](lang_dir)

    impl = f"[✅]({rel_link(impl_files[0])})" if impl_files else "—"
    test = f"[✅]({rel_link(test_files[0])})" if test_files else "—"
    return impl, test


def collect(parent: str) -> list[Path]:
    """按编号顺序收集模块目录。"""
    base = ROOT / parent
    if not base.is_dir():
        return []
    return sorted((p for p in base.iterdir() if p.is_dir()), key=lambda p: p.name)


def count_modules(modules: list[Path]) -> tuple[int, int, int, int, int]:
    """返回（模块数, Go 实现数, Go 测试数, Python 实现数, Python 测试数）。"""
    go_impl = go_test = py_impl = py_test = 0
    for module in modules:
        go_dir = module / "go"
        py_dir = module / "python"
        if go_dir.is_dir() and LANG_RULES["Go"]["impl"](go_dir):
            go_impl += 1
        if go_dir.is_dir() and LANG_RULES["Go"]["test"](go_dir):
            go_test += 1
        if py_dir.is_dir() and LANG_RULES["Python"]["impl"](py_dir):
            py_impl += 1
        if py_dir.is_dir() and LANG_RULES["Python"]["test"](py_dir):
            py_test += 1
    return len(modules), go_impl, go_test, py_impl, py_test


def module_row(module: Path, indent: str) -> str:
    """生成模块在矩阵中的一行。"""
    title = module_title(module)
    name = module.name
    readme = module / "README.md"
    label = f"[{title}]({rel_link(readme)})" if readme.is_file() else name
    go_impl, go_test = language_cells(module, "Go")
    py_impl, py_test = language_cells(module, "Python")
    return f"| {indent}{label} | {go_impl} | {go_test} | {py_impl} | {py_test} |"


def render() -> str:
    """渲染完整矩阵文档。"""
    topics = collect("topics")
    apps = collect("applications")
    topic_counts = count_modules(topics)
    app_counts = count_modules(apps)

    lines = [
        "# 覆盖矩阵",
        "",
        "> 本文件由 `scripts/gen_matrix.py` 自动生成，请勿手动编辑。",
        "> 更新实现或测试后执行 `python scripts/gen_matrix.py`。",
        "",
        "## 总览",
        "",
        "| 分类 | 模块数 | Go 实现 | Go 测试 | Python 实现 | Python 测试 |",
        "| --- | ---: | ---: | ---: | ---: | ---: |",
        "| 学习主题 | {0} | {1} | {2} | {3} | {4} |".format(*topic_counts),
        "| 应用案例 | {0} | {1} | {2} | {3} | {4} |".format(*app_counts),
        "",
        "## 学习主题",
        "",
        "| 主题 | Go 实现 | Go 测试 | Python 实现 | Python 测试 |",
        "| --- | :---: | :---: | :---: | :---: |",
    ]
    lines.extend(module_row(module, "") for module in topics)

    lines.extend(
        [
            "",
            "## 应用案例",
            "",
            "| 应用 | Go 实现 | Go 测试 | Python 实现 | Python 测试 |",
            "| --- | :---: | :---: | :---: | :---: |",
        ]
    )
    lines.extend(module_row(module, "") for module in apps)

    lines.extend(
        [
            "",
            "## 说明",
            "",
            "- ✅ 表示对应文件存在，点击可跳转；— 表示该语言尚未实现。",
            "- 规格允许同一算法只提供一种语言实现，因此矩阵存在空缺是正常的。",
            "- 所有已存在的实现都必须带测试，否则 `scripts/check_structure.py` 会失败。",
            "",
        ]
    )
    return "\n".join(lines)


def main(argv: list[str]) -> int:
    content = render()
    check_only = "--check" in argv

    if check_only:
        if not MATRIX_PATH.is_file():
            print(
                "docs/matrix.md 不存在，请先执行 python scripts/gen_matrix.py",
                file=sys.stderr,
            )
            return 1
        current = MATRIX_PATH.read_text(encoding="utf-8")
        if current != content:
            print(
                "docs/matrix.md 已过期，请执行 python scripts/gen_matrix.py 后提交更新。",
                file=sys.stderr,
            )
            return 1
        print("覆盖矩阵已是最新。")
        return 0

    MATRIX_PATH.parent.mkdir(parents=True, exist_ok=True)
    MATRIX_PATH.write_text(content, encoding="utf-8")
    print(f"已生成 {MATRIX_PATH.relative_to(ROOT).as_posix()}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv[1:]))
