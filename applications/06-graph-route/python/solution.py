"""06 路线规划 —— Python 实现。"""

from __future__ import annotations

import heapq
import math


class Router:
    """使用邻接表存储路线图。"""

    def __init__(self, nodes: int) -> None:
        self.nodes = max(0, nodes)
        self._adj: list[list[tuple[int, int]]] = [[] for _ in range(self.nodes)]

    def add_route(self, source: int, target: int, distance: int) -> None:
        """添加一条有向路线；地点越界时忽略。"""
        if self._valid(source) and self._valid(target):
            self._adj[source].append((target, distance))

    def _valid(self, node: int) -> bool:
        """判断地点下标是否合法。"""
        return 0 <= node < self.nodes

    def shortest_path(self, source: int, target: int) -> tuple[list[int], int] | None:
        """使用 Dijkstra 返回 (最短路径, 总距离)；不可达返回 None。"""
        if not self._valid(source) or not self._valid(target):
            return None

        dist = [math.inf] * self.nodes
        prev = [-1] * self.nodes
        dist[source] = 0
        queue: list[tuple[int, int]] = [(0, source)]

        while queue:
            distance, node = heapq.heappop(queue)
            if distance > dist[node]:
                continue
            if node == target:
                break
            for neighbor, weight in self._adj[node]:
                candidate = distance + weight
                if candidate < dist[neighbor]:
                    dist[neighbor] = candidate
                    prev[neighbor] = node
                    # 距离相同时按节点编号排序，保证结果确定
                    heapq.heappush(queue, (candidate, neighbor))

        if math.isinf(dist[target]):
            return None
        return self._rebuild_path(prev, source, target), dist[target]

    def fewest_hops(self, source: int, target: int) -> list[int] | None:
        """使用 BFS 返回经过地点最少的路径；不可达返回 None。"""
        if not self._valid(source) or not self._valid(target):
            return None

        prev = [-1] * self.nodes
        visited = [False] * self.nodes
        visited[source] = True
        queue = [source]

        for node in queue:
            if node == target:
                return self._rebuild_path(prev, source, target)
            for neighbor, _ in self._adj[node]:
                if visited[neighbor]:
                    continue
                visited[neighbor] = True
                prev[neighbor] = node
                queue.append(neighbor)
        return None

    def reachable(self, source: int, target: int) -> bool:
        """判断从 source 是否能到达 target。"""
        return self.fewest_hops(source, target) is not None

    def _rebuild_path(self, prev: list[int], source: int, target: int) -> list[int]:
        """根据前驱数组重建路径。"""
        path: list[int] = []
        node = target
        while node != -1:
            path.append(node)
            if node == source:
                break
            node = prev[node]
        path.reverse()
        return path
