# 06 路线规划

导航系统通常同时关心「距离最短」与「经过路口最少」，前者用 Dijkstra，后者用 BFS。

## 实现内容

| 方法 | 算法 | 时间复杂度 |
| --- | --- | --- |
| `ShortestPath` / `shortest_path` | Dijkstra + 小顶堆 | O((V + E) log V) |
| `FewestHops` / `fewest_hops` | BFS | O(V + E) |
| `Reachable` / `reachable` | BFS | O(V + E) |

## 设计要点

- 路线图使用邻接表，`add_route` 添加**有向**边，需要双向时调用两次；
- Dijkstra 的堆元素按 `(距离, 节点)` 排序，距离相同时结果仍然确定；
- 弹出节点时跳过过期记录，避免重复松弛；
- 路径通过前驱数组回溯，再反转得到起点到终点的顺序；
- 非法节点与不可达情况返回 `false` / `None`，不返回残缺路径。

## 示例

```text
0 --10--> 1 --2--> 3
 \         ^
  \--1--> 2 --20--> 4 --1--> 3
```

- 距离最短：0 → 2 → 1 → 3，总距离 4；
- 跳数最少：0 → 1 → 3，共 2 跳。

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 相关主题

- [07 图](../../topics/07-graph/README.md)
- [06 堆](../../topics/06-heap/README.md)
