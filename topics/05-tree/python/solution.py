"""05 树 —— Python 实现。"""

from __future__ import annotations

from collections import deque


class TreeNode:
    """二叉树节点。"""

    def __init__(
        self,
        val: int = 0,
        left: "TreeNode | None" = None,
        right: "TreeNode | None" = None,
    ) -> None:
        self.val = val
        self.left = left
        self.right = right


def insert_bst(root: TreeNode | None, value: int) -> TreeNode:
    """向二叉搜索树插入值，重复值不插入。"""
    if root is None:
        return TreeNode(value)
    if value < root.val:
        root.left = insert_bst(root.left, value)
    elif value > root.val:
        root.right = insert_bst(root.right, value)
    return root


def build_bst(values: list[int]) -> TreeNode | None:
    """根据列表构建二叉搜索树。"""
    root = None
    for value in values:
        root = insert_bst(root, value)
    return root


def search_bst(root: TreeNode | None, value: int) -> TreeNode | None:
    """在二叉搜索树中查找值。"""
    while root is not None:
        if value == root.val:
            return root
        root = root.left if value < root.val else root.right
    return None


def inorder(root: TreeNode | None) -> list[int]:
    """中序遍历。"""
    values: list[int] = []
    stack: list[TreeNode] = []
    node = root
    while node is not None or stack:
        while node is not None:
            stack.append(node)
            node = node.left
        node = stack.pop()
        values.append(node.val)
        node = node.right
    return values


def preorder(root: TreeNode | None) -> list[int]:
    """前序遍历。"""
    if root is None:
        return []
    return [root.val, *preorder(root.left), *preorder(root.right)]


def postorder(root: TreeNode | None) -> list[int]:
    """后序遍历。"""
    if root is None:
        return []
    return [*postorder(root.left), *postorder(root.right), root.val]


def level_order(root: TreeNode | None) -> list[list[int]]:
    """层序遍历，每层一个列表。"""
    if root is None:
        return []
    ans: list[list[int]] = []
    queue: deque[TreeNode] = deque([root])
    while queue:
        level = []
        for _ in range(len(queue)):
            node = queue.popleft()
            level.append(node.val)
            if node.left is not None:
                queue.append(node.left)
            if node.right is not None:
                queue.append(node.right)
        ans.append(level)
    return ans


def max_depth(root: TreeNode | None) -> int:
    """二叉树最大深度。"""
    if root is None:
        return 0
    return 1 + max(max_depth(root.left), max_depth(root.right))


def is_valid_bst(root: TreeNode | None) -> bool:
    """判断是否为合法二叉搜索树。"""

    def validate(node: TreeNode | None, low: float, high: float) -> bool:
        if node is None:
            return True
        if not low < node.val < high:
            return False
        return validate(node.left, low, node.val) and validate(node.right, node.val, high)

    return validate(root, float("-inf"), float("inf"))
