# shared

跨主题共享的测试数据与工具。

## 目录

```text
shared/
├── python/
│   └── solution_loader.py   # 按路径加载各模块 solution.py，避免模块名冲突
└── testdata/
    └── sort_cases.json      # 排序测试共享用例（Go / Python 共用）
```

## 使用约定

- 共享数据使用 JSON，字段命名保持 `snake_case`；
- Go 测试通过相对路径 `../../../shared/testdata/...` 读取；
- Python 测试通过 `Path(__file__).parents[3] / "shared" / ...` 读取；
- `shared/` 中不放置业务实现，保持“数据 + 测试工具”
