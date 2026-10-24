# 07 图

图用于描述对象之间的多对多关系。本主题覆盖邻接表、遍历、最短路径与拓扑排序。

## 实现内容

| 功能 | 核心思路 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- | --- |
| BFS | 队列 + visited | O(V+E) | O(V) |
| DFS | 递归 + visited | O(V+E) | O(V) |
| Dijkstra | 优先队列 + 松弛 | O((V+E) log V) | O(V+E) |
| 路径重建 | 前驱数组回溯 | O(V) | O(V) |
| 拓扑排序 | Kahn 入度法 | O(V+E) | O(V+E) |

## 实现

- Go：[solution.go](go/solution.go) · [solution_test.go](go/solution_test.go)
- Python：[solution.py](python/solution.py) · [test_solution.py](python/test_solution.py)

## 关键点

- 邻接表比邻接矩阵更省空间，适合稀疏图；
- Dijkstra 只适用于非负边权，负权需使用 Bellman-Ford；
- 优先队列中可能残留过期记录，弹出时比较当前距离即可跳过；
- 拓扑排序结果长度小于节点数说明存在环。

## 相关应用

- [06 图路径规划](../../applications/06-graph-route/README.md)

## 相关主题

- [03 栈与队列](../03-stack-queue/README.md)
- [06 堆](../06-heap/README.md)
- [11 贪心](../11-greedy/README.md)
