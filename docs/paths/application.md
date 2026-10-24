# 应用路径

1. [限流器](../../applications/01-rate-limiter/README.md)
   - 依赖：时间窗口、计数、队列
   - 核心：令牌桶与滑动窗口的实现、时间注入与并发保护。
2. [LRU 缓存](../../applications/02-lru-cache/README.md)
   - 依赖：[哈希表](../../topics/04-hash-table/README.md)、[链表](../../topics/02-linked-list/README.md)
   - 核心：哈希表 + 双向链表实现 O(1) 读写。
3. [TopK 排行](../../applications/03-topk-ranking/README.md)
   - 依赖：[堆](../../topics/06-heap/README.md)
   - 核心：小顶堆维护 TopK，避免全量排序。
4. [任务调度](../../applications/04-scheduler/README.md)
   - 依赖：[堆](../../topics/06-heap/README.md)、[贪心](../../topics/11-greedy/README.md)
   - 核心：优先级队列、分布式 ID。
5. [文本检索](../../applications/05-text-search/README.md)
   - 依赖：[字符串](../../topics/13-string/README.md)
   - 核心：倒排索引与前缀树（Trie）。
6. [图路径规划](../../applications/06-graph-route/README.md)
   - 依赖：[图](../../topics/07-graph/README.md)
   - 核心：Dijkstra 与路径重建。
7. [去重与布隆过滤](../../applications/07-dedup-bloom/README.md)
   - 核心：概率数据结构的误判率与空间权衡。
8. [流式统计](../../applications/08-stream-stats/README.md)
   - 核心：Count-Min Sketch、滑动平均、蓄水池抽样。
9. [缓存淘汰策略](../../applications/09-cache-eviction/README.md)
   - 核心：FIFO、LFU 与 LRU 的取舍。

## 工程关注点

- 并发安全：锁、原子操作、协程退出；
- 可测试性：注入时钟、固定随机种子；
- 可观测性：命中率、淘汰次数等指标；
- 复杂度：把数据结构选择与访问模式匹配。
