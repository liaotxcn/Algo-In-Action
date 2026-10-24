# 02 LRU 缓存

LRU（Least Recently Used）在容量不足时淘汰最久未被访问的数据。它被广泛用于数据库缓冲池、HTTP 缓存与本地热数据缓存。

## 数据结构

```text
map[key] -> 链表节点指针
链表头部：最近使用
链表尾部：最久未使用
```

- `Get`：命中后把节点移动到头部，O(1)；
- `Put`：键存在则更新并移到头部；键不存在且已满则删除尾部节点后插入头部，O(1)；
- 使用互斥锁保证并发安全。

## 实现

| 语言 | 底层容器 |
| --- | --- |
| Go | `container/list` + `map` |
| Python | `collections.OrderedDict` |

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 接口

| 操作 | Go | Python |
| --- | --- | --- |
| 读取 | `Get(key) (V, bool)` | `get(key, default=None)` |
| 写入 | `Put(key, value)` | `put(key, value)` |
| 长度 | `Len() int` | `len(cache)` |
| 统计 | `Stats() (hits, misses)` | `cache.stats` |

## 测试覆盖

- 基本读写与淘汰顺序；
- 更新已有键后刷新热度；
- 容量为 0 的边界行为；
- 命中与未命中计数。

## 相关主题

- [02 链表](../../topics/02-linked-list/README.md)
- [04 哈希表](../../topics/04-hash-table/README.md)
- [09 缓存淘汰](../09-cache-eviction/README.md)
