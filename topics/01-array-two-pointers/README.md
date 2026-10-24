# 01 数组与双指针

数组是最基础的线性结构。本主题围绕下标操作、双指针、前缀和与原地修改展开，并补充矩阵类高频题。

## 题目与实现

| 题目 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| 轮转数组 | 三次反转 | O(n) | O(1) |
| 盛最多水的容器 | 左右双指针，移动较短边 | O(n) | O(1) |
| 三数之和 | 排序 + 双指针 + 去重 | O(n²) | O(1) |
| 移动零 | 读写指针 | O(n) | O(1) |
| 接雨水 | 左右双指针维护两侧最大值 | O(n) | O(1) |
| 合并区间 | 排序后合并 | O(n log n) | O(n) |
| 除自身以外数组的乘积 | 前后缀乘积 | O(n) | O(n) |
| 最大子数组和 | 前缀和 + 最小前缀 | O(n) | O(1) |
| 螺旋矩阵 | 按层模拟边界 | O(mn) | O(mn) |
| 矩阵置零 | 复用第一行/第一列做标记 | O(mn) | O(1) |
| 旋转图像 | 转置 + 反转每行 | O(n²) | O(1) |
| 搜索二维矩阵 II | 从右上角开始收缩 | O(m+n) | O(1) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)
- 共享用例：[cases/cases.json](cases/cases.json)（Go / Python 测试共同读取）

## 关键点

- 双指针题先确定「指针移动后搜索空间一定缩小」的不变量；
- 原地修改题注意读写指针的先后顺序；
- 区间/矩阵题先处理空输入与单行单列边界。

## 相关应用

- [03 TopK 排行](../../applications/03-topk-ranking/README.md)
- [05 文本检索](../../applications/05-text-search/README.md)

## 相关主题

- [04 哈希表](../04-hash-table/README.md)：前缀和与哈希结合
- [09 搜索](../09-searching/README.md)：有序数组上的二分
