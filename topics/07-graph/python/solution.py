"""07 图 —— Python 实现。"""

from __future__ import annotations

import heapq
from collections import deque


class Graph:
    """邻接表图。"""

    def __init__(self, nodes: int) -> None:
        self.nodes = nodes
        self.adj: list[list[tuple[int, int]]] = [[] for _ in range(nodes)]

    def add_edge(self, src: int, dst: int, weight: int = 1) -> None:
        """添加有向边。"""
        if 0 <= src < self.nodes and 0 <= dst < self.nodes:
            self.adj[src].append((dst, weight))

    def add_undirected_edge(self, a: int, b: int, weight: int = 1) -> None:
        """添加无向边。"""
        self.add_edge(a, b, weight)
        self.add_edge(b, a, weight)

    def bfs(self, start: int) -> list[int]:
        """广度优先遍历。"""
        if not 0 <= start < self.nodes:
            return []
        visited = [False] * self.nodes
        visited[start] = True
        queue: deque[int] = deque([start])
        order: list[int] = []
        while queue:
            node = queue.popleft()
            order.append(node)
            for neighbor, _ in self.adj[node]:
                if not visited[neighbor]:
                    visited[neighbor] = True
                    queue.append(neighbor)
        return order

    def dfs(self, start: int) -> list[int]:
        """深度优先遍历。"""
        if not 0 <= start < self.nodes:
            return []
        visited = [False] * self.nodes
        order: list[int] = []

        def visit(node: int) -> None:
            visited[node] = True
            order.append(node)
            for neighbor, _ in self.adj[node]:
                if not visited[neighbor]:
                    visit(neighbor)

        visit(start)
        return order

    def dijkstra(self, start: int) -> tuple[list[int], list[int]]:
        """返回最短距离与前驱数组，不可达距离为 math.inf。"""
        dist = [float("inf")] * self.nodes
        prev = [-1] * self.nodes
        if not 0 <= start < self.nodes:
            return dist, prev

        dist[start] = 0
        pq: list[tuple[int, int]] = [(0, start)]
        while pq:
            current, node = heapq.heappop(pq)
            if current > dist[node]:
                continue
            for neighbor, weight in self.adj[node]:
                candidate = current + weight
                if candidate < dist[neighbor]:
                    dist[neighbor] = candidate
                    prev[neighbor] = node
                    heapq.heappush(pq, (candidate, neighbor))
        return dist, prev

    def build_path(self, prev: list[int], target: int) -> list[int]:
        """根据前驱数组重建路径。"""
        if not 0 <= target < len(prev):
            return []
        path = []
        node = target
        while node != -1:
            path.append(node)
            node = prev[node]
        return path[::-1]

    def topological_sort(self) -> list[int] | None:
        """Kahn 拓扑排序；有环返回 None。"""
        indegree = [0] * self.nodes
        for edges in self.adj:
            for neighbor, _ in edges:
                indegree[neighbor] += 1

        queue: deque[int] = deque(
            node for node, degree in enumerate(indegree) if degree == 0
        )
        order: list[int] = []
        while queue:
            node = queue.popleft()
            order.append(node)
            for neighbor, _ in self.adj[node]:
                indegree[neighbor] -= 1
                if indegree[neighbor] == 0:
                    queue.append(neighbor)
        return order if len(order) == self.nodes else None
