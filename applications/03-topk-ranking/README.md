# 03 TopK 排行

从大量数据中找出最大的 K 个元素或出现频率最高的 K 个词，是榜单、热搜、推荐系统的常见需求。

## 算法

- 维护一个大小为 K 的**小顶堆**，堆顶是当前 TopK 中最差的元素；
- 新元素优于堆顶时替换堆顶，最终堆内即为 TopK；
- 时间复杂度 O(n log k)，空间复杂度 O(k)，适合 n 远大于 k 的场景。

## 实现内容

| 函数 | 说明 | 排序规则 |
| --- | --- | --- |
| `TopK` / `top_k` | 最大的 K 个数 | 数值降序 |
| `TopKFrequent` / `top_k_frequent` | 频率最高的 K 个词 | 次数降序，次数相同按字典序升序 |

Go 使用 `container/heap`，Python 使用 `heapq`。Python 通过 `_WordFreq.__lt__` 定义「更差者更小」，保证堆顶始终是当前最应淘汰的候选。

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 测试覆盖

- 普通数组与含重复元素的数组；
- `k = 0`、空数组、`k` 超过数组长度；
- 词频排行与同频次字典序 tie-break。

## 相关主题

- [06 堆](../../topics/06-heap/README.md)
- [08 排序](../../topics/08-sorting/README.md)
- [08 流式统计](../08-stream-stats/README.md)
