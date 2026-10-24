# 03 栈与队列

栈（LIFO）和队列（FIFO）是两种最常用的受限线性结构，也是单调栈/单调队列等技巧的基础。

## 题目与实现

| 题目 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 有效的括号 | 栈匹配 | O(n) | O(n) |
| 最小栈 | 辅助栈记录历史最小值 | O(1) | O(n) |
| 用栈实现队列 | 入栈/出栈双栈转移 | 均摊 O(1) | O(n) |
| 滑动窗口最大值 | 单调递减队列 | O(n) | O(k) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 最小栈的辅助栈保存「当前栈状态下的最小值」；
- 双栈队列只在出栈为空时转移，保证均摊 O(1)；
- 单调队列存下标而不是值，便于判断是否滑出窗口。

## 相关应用

- [01 限流器](../../applications/01-rate-limiter/README.md)：滑动窗口
- [04 任务调度](../../applications/04-scheduler/README.md)：优先队列

## 相关主题

- [02 链表](../02-linked-list/README.md)
- [06 堆](../06-heap/README.md)
