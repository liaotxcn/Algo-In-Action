#!/usr/bin/env bash
# 本地测试聚合入口：结构校验 -> Go 测试 -> Python 测试
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

echo "==> 校验目录结构"
python scripts/check_structure.py

echo "==> 运行 Go 测试"
go test ./...

echo "==> 运行 Python 测试"
python -m pytest -q

echo "全部检查通过。"
