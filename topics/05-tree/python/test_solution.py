"""05 树的 pytest 测试。"""

from solution_loader import load_solution

solution = load_solution(__file__)


def test_insert_and_search_bst() -> None:
    root = solution.build_bst([5, 3, 8, 1, 4, 7, 9, 3])
    assert solution.inorder(root) == [1, 3, 4, 5, 7, 8, 9]
    assert solution.search_bst(root, 4).val == 4
    assert solution.search_bst(root, 6) is None


def test_traversals() -> None:
    root = solution.build_bst([4, 2, 6, 1, 3, 5, 7])
    assert solution.preorder(root) == [4, 2, 1, 3, 6, 5, 7]
    assert solution.inorder(root) == [1, 2, 3, 4, 5, 6, 7]
    assert solution.postorder(root) == [1, 3, 2, 5, 7, 6, 4]


def test_level_order() -> None:
    root = solution.TreeNode(
        3,
        solution.TreeNode(9),
        solution.TreeNode(20, solution.TreeNode(15), solution.TreeNode(7)),
    )
    assert solution.level_order(root) == [[3], [9, 20], [15, 7]]
    assert solution.level_order(None) == []


def test_max_depth() -> None:
    root = solution.TreeNode(
        3,
        solution.TreeNode(9),
        solution.TreeNode(20, solution.TreeNode(15), solution.TreeNode(7)),
    )
    assert solution.max_depth(root) == 3
    assert solution.max_depth(None) == 0


def test_is_valid_bst() -> None:
    assert solution.is_valid_bst(solution.build_bst([2, 1, 3])) is True
    invalid = solution.TreeNode(
        5,
        solution.TreeNode(1),
        solution.TreeNode(4, solution.TreeNode(3), solution.TreeNode(6)),
    )
    assert solution.is_valid_bst(invalid) is False
