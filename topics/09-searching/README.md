# 09 搜索

搜索用于在数据集中定位目标。有序数组上的二分查找是最常用的 O(log n) 手段，边界变体尤其重要。

## 实现内容

| 功能 | 说明 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 二分查找 | 返回任意匹配下标 | O(log n) | O(1) |
| LowerBound | 第一个 `>= target` 的位置 | O(log n) | O(1) |
| UpperBound | 第一个 `> target` 的位置 | O(log n) | O(1) |
| SearchRange | 目标值的起止区间 | O(log n) | O(1) |
| 插值查找 | 按数值比例估计位置 | 平均 O(log log n) | O(1) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 统一使用左闭右开区间 `[left, right)` 可以减少边界错误；
- `mid = left + (right-left)/2` 避免加法溢出；
- 找第一个满足条件的位置时，把「判定」与「移动」分开；
- 插值查找只适合分布均匀的数据，最坏会退化为 O(n)。

## 相关应用

- [05 文本检索](../../applications/05-text-search/README.md)
- [03 TopK 排行](../../applications/03-topk-ranking/README.md)

## 相关主题

- [08 排序](../08-sorting/README.md)
- [01 数组与双指针](../01-array-two-pointers/README.md)
