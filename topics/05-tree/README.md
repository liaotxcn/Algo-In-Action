# 05 树

树是最常见的非线性结构。本主题从二叉树的遍历出发，延伸到二叉搜索树的插入、查找与合法性校验。

## 实现内容

| 功能 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| BST 插入/构建 | 递归比较大小 | 平均 O(log n) | O(h) |
| BST 查找 | 迭代下降 | 平均 O(log n) | O(1) |
| 前/中/后序遍历 | 递归或显式栈 | O(n) | O(h) |
| 层序遍历 | 队列 BFS | O(n) | O(n) |
| 最大深度 | 递归取左右最大值 | O(n) | O(h) |
| 验证 BST | 传递上下界 | O(n) | O(h) |

> h 为树高：平衡树 O(log n)，退化成链表时 O(n)。

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 验证 BST 不能只比较父子节点，必须传递 `(low, high)` 区间；
- 中序遍历 BST 得到升序序列；
- 层序遍历用队列，按当前层长度切分。

## 相关应用

- [04 任务调度](../../applications/04-scheduler/README.md)
- [03 TopK 排行](../../applications/03-topk-ranking/README.md)

## 相关主题

- [06 堆](../06-heap/README.md)
- [07 图](../07-graph/README.md)
