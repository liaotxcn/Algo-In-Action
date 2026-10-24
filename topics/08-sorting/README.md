# 08 排序

排序是最基础的数据预处理手段。本主题实现比较排序与非比较排序，并通过共享用例保证两种语言结果一致。

## 实现内容

| 算法 | 类别 | 平均时间复杂度 | 最坏 | 空间 | 稳定 |
| --- | --- | --- | --- | --- | --- |
| 冒泡排序 | 比较 | O(n²) | O(n²) | O(1) | 是 |
| 选择排序 | 比较 | O(n²) | O(n²) | O(1) | 否 |
| 插入排序 | 比较 | O(n²) | O(n²) | O(1) | 是 |
| 希尔排序 | 比较 | O(n^1.3) | O(n²) | O(1) | 否 |
| 归并排序 | 比较 | O(n log n) | O(n log n) | O(n) | 是 |
| 快速排序 | 比较 | O(n log n) | O(n²) | O(log n) | 否 |
| 堆排序 | 比较 | O(n log n) | O(n log n) | O(1) | 否 |
| 计数排序 | 非比较 | O(n+k) | O(n+k) | O(k) | 是 |
| 桶排序 | 非比较 | O(n+k) | O(n²) | O(n+k) | 取决于桶内排序 |
| 基数排序 | 非比较 | O(d(n+10)) | O(d(n+10)) | O(n) | 是 |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)
- 共享用例：[shared/testdata/sort_cases.json](../../shared/testdata/sort_cases.json)

> Go 版本原地排序；Python 版本返回新列表，不修改入参。

## 如何选择

- 小规模或基本有序：插入排序、冒泡排序；
- 大规模通用：快速排序、归并排序、堆排序；
- 稳定且需要链表/外部排序：归并排序；
- 整数且范围有限：计数排序；
- 分布均匀：桶排序；
- 位数少的大整数：基数排序。

## 相关应用

- [03 TopK 排行](../../applications/03-topk-ranking/README.md)
- [05 文本检索](../../applications/05-text-search/README.md)

## 相关主题

- [06 堆](../06-heap/README.md)
- [09 搜索](../09-searching/README.md)
