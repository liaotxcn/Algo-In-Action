"""06 路线规划的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def build_router() -> "solution.Router":
    router = solution.Router(5)
    router.add_route(0, 1, 10)
    router.add_route(0, 2, 1)
    router.add_route(2, 1, 1)
    router.add_route(1, 3, 2)
    router.add_route(2, 4, 20)
    router.add_route(4, 3, 1)
    return router


def test_shortest_path() -> None:
    path, distance = build_router().shortest_path(0, 3)
    assert path == [0, 2, 1, 3]
    assert distance == 4


def test_fewest_hops() -> None:
    assert build_router().fewest_hops(0, 3) == [0, 1, 3]


def test_unreachable() -> None:
    router = build_router()
    assert router.shortest_path(3, 0) is None
    assert router.fewest_hops(3, 0) is None
    assert router.reachable(0, 3)
    assert not router.reachable(3, 0)


def test_invalid_nodes() -> None:
    router = build_router()
    assert router.shortest_path(-1, 3) is None
    assert router.fewest_hops(0, 99) is None
