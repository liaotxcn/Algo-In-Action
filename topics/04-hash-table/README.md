# 04 哈希表

哈希表用空间换时间，把查找、插入、删除的期望复杂度降到 O(1)。本主题关注键的设计与「前缀和 + 哈希」的组合。

## 题目与实现

| 题目 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 两数之和 | 边遍历边记录补数 | O(n) | O(n) |
| 最长连续序列 | 哈希集合 + 只从起点扩展 | O(n) | O(n) |
| 字母异位词分组 | 排序后的字符串作为键 | O(n·k log k) | O(n·k) |
| 和为 K 的子数组 | 前缀和 + 哈希计数 | O(n) | O(n) |
| 有效的字母异位词 | 26 位计数 | O(n) | O(1) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 先想清楚「键是什么」：值、排序后的字符串、前缀和；
- 前缀和问题初始化 `{0: 1}` 表示空前缀；
- 输出分组时排序，保证测试与示例稳定。

## 相关应用

- [02 LRU 缓存](../../applications/02-lru-cache/README.md)
- [07 去重与布隆过滤](../../applications/07-dedup-bloom/README.md)
- [08 流式统计](../../applications/08-stream-stats/README.md)

## 相关主题

- [01 数组与双指针](../01-array-two-pointers/README.md)
- [13 字符串](../13-string/README.md)
