# 12 回溯

回溯是带撤销操作的深度优先搜索：在解空间树上做选择、递归、撤销选择。剪枝决定它能否在时限内跑完。

## 实现内容

| 问题 | 剪枝方式 | 时间复杂度 |
| --- | --- | --- |
| 全排列 | used 数组标记已选元素 | O(n · n!) |
| 去重全排列 | 排序后跳过同层重复元素 | O(n · n!) |
| 组合 | 剩余元素不足以填满时提前结束 | O(C(n, k) · k) |
| 子集 | 每层从 start 向后扩展 | O(n · 2^n) |
| N 皇后 | 列、主对角线、副对角线布尔数组 | O(n!) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 回溯模板

```text
backtrack(path, 选择列表):
    if 满足结束条件:
        记录结果
        return
    for 选择 in 选择列表:
        if 剪枝条件: continue
        做选择
        backtrack(path, 新的选择列表)
        撤销选择
```

## 关键点

- 组合类问题用 `start` 控制顺序，避免生成重复排列；
- 去重全排列的 `!used[i-1]` 判断表示「同一层已用过相同元素」；
- 棋盘类问题用布尔数组代替每次扫描，把合法性检查降到 O(1)；
- 剪枝条件写错通常表现为「结果缺失」而不是「结果多余」，测试要覆盖边界。

## 相关主题

- [05 树](../05-tree/README.md)
- [07 图](../07-graph/README.md)
- [10 动态规划](../10-dp/README.md)
