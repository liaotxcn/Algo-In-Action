# Algo-In-Action

[English](README.md) | 简体中文

「数据结构与算法 + 应用落地」

每个主题均遵循同一套约定：文档、可运行实现、自动化测试；主题内容与应用场景互相链接，形成从原理到应用落地的闭环

## 从哪里开始

| 入口 | 说明 |
| --- | --- |
| [导航入口](docs/index.md) | 按路径、覆盖矩阵 |
| [应用路径](docs/paths/application.md) | 限流、缓存、检索、调度等落地案例 |
| [覆盖矩阵](docs/matrix.md) | 每个主题的语言实现与测试覆盖情况 |

## 目录结构

```text
Algo-In-Action/
├── docs/          # 主题路径、覆盖矩阵
├── topics/        # 内容主体
├── applications/  # 应用落地
├── shared/        # 跨主题共享测试数据与工具
└── scripts/       # 结构校验、矩阵生成、测试聚合
```

## 本地验证

环境要求：Go 1.26+、Python 3.11+、pytest

```bash
python -m pip install pytest
bash scripts/run_all_tests.sh
```

也可以单独执行：

```bash
python scripts/check_structure.py   # 校验目录约定
python scripts/gen_matrix.py        # 重新生成 docs/matrix.md
go test ./...                       # Go 测试
python -m pytest -q                 # Python 测试
```

## 内容约定

- 主题实现：`topics/<nn-topic>/go/solution.go` 与 `python/solution.py`
- 主题测试：`solution_test.go` 与 `test_solution.py`
- 应用实现：`applications/<nn-app>/` 下按语言组织
- 新增或调整模块后，需运行结构校验与矩阵生成
