"""07 图的 pytest 测试。"""

import math

from solution_loader import load_solution

solution = load_solution(__file__)


def build_weighted_graph():
    graph = solution.Graph(6)
    graph.add_undirected_edge(0, 1, 4)
    graph.add_undirected_edge(0, 2, 2)
    graph.add_undirected_edge(1, 2, 1)
    graph.add_undirected_edge(1, 3, 5)
    graph.add_undirected_edge(2, 3, 8)
    graph.add_undirected_edge(2, 4, 10)
    graph.add_undirected_edge(3, 4, 2)
    graph.add_undirected_edge(3, 5, 6)
    graph.add_undirected_edge(4, 5, 3)
    return graph


def test_bfs_and_dfs() -> None:
    graph = solution.Graph(4)
    graph.add_edge(0, 1)
    graph.add_edge(0, 2)
    graph.add_edge(1, 3)
    graph.add_edge(2, 3)
    assert graph.bfs(0) == [0, 1, 2, 3]
    assert graph.dfs(0) == [0, 1, 3, 2]


def test_dijkstra_and_path() -> None:
    graph = build_weighted_graph()
    dist, prev = graph.dijkstra(0)
    assert dist == [0, 3, 2, 8, 10, 13]
    assert graph.build_path(prev, 5) == [0, 2, 1, 3, 4, 5]


def test_dijkstra_unreachable() -> None:
    graph = solution.Graph(3)
    graph.add_edge(0, 1, 5)
    dist, prev = graph.dijkstra(0)
    assert dist[2] == math.inf
    assert graph.build_path(prev, 2) == [2]


def test_topological_sort() -> None:
    graph = solution.Graph(5)
    graph.add_edge(0, 1)
    graph.add_edge(0, 2)
    graph.add_edge(1, 3)
    graph.add_edge(2, 3)
    graph.add_edge(3, 4)
    assert graph.topological_sort() == [0, 1, 2, 3, 4]

    cyclic = solution.Graph(2)
    cyclic.add_edge(0, 1)
    cyclic.add_edge(1, 0)
    assert cyclic.topological_sort() is None
