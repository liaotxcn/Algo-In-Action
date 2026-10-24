# 10 动态规划

动态规划通过保存子问题结果避免重复计算。关键是定义状态、写出转移方程，并确定遍历顺序。

## 实现内容

| 问题 | 状态定义 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 斐波那契 | 前两项滚动 | O(n) | O(1) |
| 爬楼梯 | `f(n)=f(n-1)+f(n-2)` | O(n) | O(1) |
| 0-1 背包 | `dp[c]` 容量 c 的最大价值 | O(n·C) | O(C) |
| 最长递增子序列 | 贪心 + 二分维护 tails | O(n log n) | O(n) |
| 编辑距离 | 二维 DP 压缩为一维 | O(m·n) | O(n) |
| 零钱兑换 | `dp[amount]` 最少硬币数 | O(n·amount) | O(amount) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 0-1 背包空间压缩后必须**逆序**遍历容量；
- 完全背包（零钱兑换）使用**正序**；
- 编辑距离注意空串边界；
- 无法到达的状态用哨兵值表示，最后再转换。

## 相关应用

- [08 流式统计](../../applications/08-stream-stats/README.md)
- [05 文本检索](../../applications/05-text-search/README.md)

## 相关主题

- [09 搜索](../09-searching/README.md)
- [11 贪心](../11-greedy/README.md)
- [12 回溯](../12-backtracking/README.md)
