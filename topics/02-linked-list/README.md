# 02 链表

链表通过指针把离散节点串联起来，考察点集中在指针操作、边界处理和空间复杂度优化。

## 题目与实现

| 题目 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 反转链表 | 三指针迭代 | O(n) | O(1) |
| 环形链表 | 快慢指针 | O(n) | O(1) |
| 相交链表 | 双指针走完两条链表 | O(m+n) | O(1) |
| 链表的中间节点 | 快慢指针 | O(n) | O(1) |
| 回文链表 | 找中点 + 反转后半段 | O(n) | O(1) |
| 合并两个有序链表 | 哨兵节点 + 尾指针 | O(m+n) | O(1) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 使用哨兵节点（dummy）简化头节点处理；
- 修改 `next` 前先保存后继，避免断链；
- 快慢指针的循环条件通常是 `fast != nil && fast.Next != nil`；
- 相交链表的核心是让两个指针走相同的总路程。

## 相关应用

- [02 LRU 缓存](../../applications/02-lru-cache/README.md)：哈希表 + 双向链表
- [09 缓存淘汰策略](../../applications/09-cache-eviction/README.md)

## 相关主题

- [04 哈希表](../04-hash-table/README.md)
- [03 栈与队列](../03-stack-queue/README.md)
