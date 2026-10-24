# 11 贪心

贪心算法每一步都选择当前看起来最优的方案，并要求该选择不会破坏全局最优。难点不在实现，而在证明贪心策略正确。

## 实现内容

| 问题 | 贪心策略 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 活动选择 | 按结束时间最早选择 | O(n log n) | O(n) |
| 分数背包 | 按单位价值从高到低装 | O(n log n) | O(n) |
| 跳跃游戏 | 维护最远可达位置 | O(n) | O(1) |
| 最少跳跃次数 | 按层扩展当前可达边界 | O(n) | O(1) |
| 分发饼干 | 小饼干优先满足小胃口 | O(n log n) | O(1) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 贪心题先找「排序关键字」或「局部最优指标」；
- 活动选择按结束时间排序，而不是开始时间；
- 分数背包可以拆分物品，0-1 背包不可以，因此后者需要 DP；
- 最少跳跃次数使用 BFS 分层思想实现 O(n)。

## 相关应用

- [04 任务调度](../../applications/04-scheduler/README.md)
- [03 TopK 排行](../../applications/03-topk-ranking/README.md)

## 相关主题

- [10 动态规划](../10-dp/README.md)
- [08 排序](../08-sorting/README.md)
