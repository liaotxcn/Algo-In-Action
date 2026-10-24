# 06 堆

堆是一棵完全二叉树，适合维护动态集合的最值。核心操作是向上/向下调整，代价 O(log n)。

## 实现内容

| 功能 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 大顶堆 Push/Pop | 上浮/下沉 | O(log n) | O(1) |
| 堆排序 | 建堆 + 反复取堆顶 | O(n log n) | O(1) |
| TopK | 大小为 k 的小顶堆 | O(n log k) | O(k) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 下标从 0 开始时：父节点 `(i-1)/2`，左孩子 `2i+1`，右孩子 `2i+2`；
- 建堆从最后一个非叶子节点 `n/2-1` 开始；
- TopK 用小顶堆：堆顶是当前候选集中最小的，便于淘汰；
- Python 的 `heapq` 是小顶堆，做大顶堆时存入负数。

## 相关应用

- [03 TopK 排行](../../applications/03-topk-ranking/README.md)
- [04 任务调度](../../applications/04-scheduler/README.md)

## 相关主题

- [08 排序](../08-sorting/README.md)
- [05 树](../05-tree/README.md)
