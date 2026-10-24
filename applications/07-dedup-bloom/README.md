# 07 去重过滤器

布隆过滤器与布谷鸟过滤器都用少量内存判断「元素是否可能存在」，适合 URL 去重、缓存穿透防护、黑名单等场景。本仓库只使用标准库实现，不依赖 Redis 或第三方哈希库。

## 布隆过滤器

- 位数组长度 `m = -n · ln p / (ln 2)²`，哈希函数个数 `k = (m/n) · ln 2`；
- 使用双重哈希生成 k 个位下标，无需 k 个独立哈希函数；
- **没有假阴性**：判定不存在则一定不存在；
- **存在假阳性**：判定存在时可能误判；
- 不支持删除（会破坏其他元素的位）。

## 布谷鸟过滤器

- 每个元素保存 16 位指纹，放入两个候选桶之一；
- 两个桶都满时踢出一个指纹，被踢出的指纹迁移到它的备用桶；
- 超过最大踢出次数则插入失败；
- 支持删除，但同一元素反复插入/删除可能留下多个指纹。

## 复杂度

| 操作 | 布隆过滤器 | 布谷鸟过滤器 |
| --- | --- | --- |
| Add | O(k) | 均摊 O(1)，最坏 O(maxKicks) |
| MightContain | O(k) | O(bucketSize) |
| Delete | 不支持 | O(bucketSize) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 测试覆盖

- 已插入元素无假阴性；
- 误判率不超过宽松阈值（10%）；
- 布谷鸟过滤器插入、删除与计数。

## 相关主题

- [04 哈希表](../../topics/04-hash-table/README.md)
- [13 字符串](../../topics/13-string/README.md)
