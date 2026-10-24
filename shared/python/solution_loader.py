"""跨模块加载 solution.py 的小工具。

每个模块的 python/ 目录都有一个名为 solution.py 的文件，直接 `import solution`
会因为模块名重复而互相覆盖。这里按文件路径显式加载，保证测试彼此隔离。
"""

from __future__ import annotations

import importlib.util
import sys
from pathlib import Path
from types import ModuleType


def load_solution(test_file: str | Path) -> ModuleType:
    """加载与测试文件同目录的 solution.py。"""
    test_path = Path(test_file).resolve()
    solution_path = test_path.with_name("solution.py")
    if not solution_path.is_file():
        raise FileNotFoundError(f"未找到实现文件: {solution_path}")

    # 用路径片段拼出唯一模块名，避免不同模块之间污染 sys.modules
    parts = [p.replace("-", "_") for p in solution_path.parts[-4:-1]]
    module_name = "algohub_" + "_".join(parts)

    spec = importlib.util.spec_from_file_location(module_name, solution_path)
    if spec is None or spec.loader is None:
        raise ImportError(f"无法加载实现文件: {solution_path}")

    module = importlib.util.module_from_spec(spec)
    # dataclass 等机制要求模块在执行前已注册到 sys.modules
    sys.modules[module_name] = module
    try:
        spec.loader.exec_module(module)
    except BaseException:
        sys.modules.pop(module_name, None)
        raise
    return module
