# Algo-In-Action

English | [简体中文](README.zh-CN.md)

“Data Structures & Algorithms + Real-World Applications”

Every topic follows the same conventions: documentation, runnable implementations, and automated tests. Topics and application scenarios are cross-linked, forming a closed loop from principles to real-world application.

## Getting Started

| Entry | Description |
| --- | --- |
| [Documentation Home](docs/index.md) | Browse by learning path and coverage matrix |
| [Application Path](docs/paths/application.md) | Practical cases such as rate limiting, caching, search, and scheduling |
| [Coverage Matrix](docs/matrix.md) | Language implementations and test coverage for each topic |

## Directory Structure

```text
Algo-In-Action/
├── docs/          # Topic paths, coverage matrix
├── topics/        # Core learning content
├── applications/  # Real-world applications
├── shared/        # Shared test data and utilities across topics
└── scripts/       # Structure validation, matrix generation, and test aggregation
```

## Local Verification

Requirements: Go 1.26+, Python 3.11+, pytest

```bash
python -m pip install pytest
bash scripts/run_all_tests.sh
```

You can also run individual checks:

```bash
python scripts/check_structure.py   # Validate directory conventions
python scripts/gen_matrix.py        # Regenerate docs/matrix.md
go test ./...                       # Run Go tests
python -m pytest -q                 # Run Python tests
```

## Content Conventions

- Topic implementations: `topics/<nn-topic>/go/solution.go` and `python/solution.py`
- Topic tests: `solution_test.go` and `test_solution.py`
- Application implementations: organized by language under `applications/<nn-app>/`
- Run structure validation and matrix generation after adding or adjusting modules
