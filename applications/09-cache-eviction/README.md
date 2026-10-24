# 09 缓存淘汰策略

当缓存容量有限时必须选择淘汰对象。本应用实现两种经典策略，并与 [LRU 缓存](../02-lru-cache/README.md) 对照。

## FIFO（先进先出）

- 新数据追加到队尾，容量不足时淘汰队头；
- 读取不改变顺序，实现简单；
- 缺点：热点数据可能因写入早而被误淘汰。

## LFU（最不经常使用）

- 每个条目维护访问频次；
- 按频次分桶，每个桶内部按最近使用顺序排列；
- 淘汰最低频次桶中最久未使用的条目；
- 频次相同时退化为 LRU，避免「旧热点」长期占用缓存。

## 复杂度

| 操作 | FIFO | LFU |
| --- | --- | --- |
| Get | O(1) | O(1) |
| Put | O(1) | O(1) |
| 空间 | O(capacity) | O(capacity) |

## 实现

| 语言 | FIFO | LFU |
| --- | --- | --- |
| Go | `container/list` + `map` | 频次桶 `map[int]*list.List` |
| Python | `OrderedDict` | 频次桶 `dict[int, OrderedDict]` |

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 测试覆盖

- FIFO 按写入顺序淘汰，更新不改变顺序；
- LFU 淘汰最低频次条目；
- 同频次时淘汰最久未使用条目；
- 容量为 0 的边界行为。

## 相关应用

- [02 LRU 缓存](../02-lru-cache/README.md)

## 相关主题

- [02 链表](../../topics/02-linked-list/README.md)
- [04 哈希表](../../topics/04-hash-table/README.md)
