# 导航

本仓库分为两条主线：

- **topics/**：数据结构/算法学习主体，覆盖 13 个经典主题。
- **applications/**：真实场景应用，覆盖 9 个应用方向。

两条主线通过每篇 README 中的「相关应用」「相关主题」互相链接。

## 覆盖矩阵

[覆盖矩阵](matrix.md) 由 `scripts/gen_matrix.py` 自动生成，可快速查看某个主题有哪些语言实现、是否带测试。

## 主题

| 编号 | 主题 | 相关应用 |
| --- | --- | --- |
| 01 | [数组与双指针](../topics/01-array-two-pointers/README.md) | 03 TopK、05 文本检索 |
| 02 | [链表](../topics/02-linked-list/README.md) | 02 LRU 缓存 |
| 03 | [栈与队列](../topics/03-stack-queue/README.md) | 04 调度器 |
| 04 | [哈希表](../topics/04-hash-table/README.md) | 02 LRU 缓存、07 去重 |
| 05 | [树](../topics/05-tree/README.md) | 05 文本检索 |
| 06 | [堆](../topics/06-heap/README.md) | 03 TopK、04 调度器 |
| 07 | [图](../topics/07-graph/README.md) | 06 路径规划 |
| 08 | [排序](../topics/08-sorting/README.md) | 03 TopK |
| 09 | [搜索](../topics/09-searching/README.md) | 05 文本检索 |
| 10 | [动态规划](../topics/10-dp/README.md) | — |
| 11 | [贪心](../topics/11-greedy/README.md) | 04 调度器 |
| 12 | [回溯](../topics/12-backtracking/README.md) | — |
| 13 | [字符串](../topics/13-string/README.md) | 05 文本检索、07 去重 |

## 应用

| 编号 | 案例 | 依赖主题 |
| --- | --- | --- |
| 01 | [限流器](../applications/01-rate-limiter/README.md) | 队列、时间窗口 |
| 02 | [LRU 缓存](../applications/02-lru-cache/README.md) | 哈希表、链表 |
| 03 | [TopK 排行](../applications/03-topk-ranking/README.md) | 堆、排序 |
| 04 | [任务调度](../applications/04-scheduler/README.md) | 堆、贪心 |
| 05 | [文本检索](../applications/05-text-search/README.md) | 字符串、搜索 |
| 06 | [图路径规划](../applications/06-graph-route/README.md) | 图 |
| 07 | [去重与布隆过滤](../applications/07-dedup-bloom/README.md) | 哈希、字符串 |
| 08 | [流式统计](../applications/08-stream-stats/README.md) | 哈希、概率数据结构 |
| 09 | [缓存淘汰策略](../applications/09-cache-eviction/README.md) | 链表、哈希、堆 |
