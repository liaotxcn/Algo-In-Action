// Package solution 提供二叉树与二叉搜索树相关算法的 Go 实现。
package solution

import "math"

// TreeNode 是二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InsertBST 向二叉搜索树插入值，返回根节点；重复值不插入。
// 平均时间复杂度 O(log n)，最坏 O(n)。
func InsertBST(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: value}
	}
	switch {
	case value < root.Val:
		root.Left = InsertBST(root.Left, value)
	case value > root.Val:
		root.Right = InsertBST(root.Right, value)
	}
	return root
}

// BuildBST 根据切片构建二叉搜索树。
func BuildBST(values []int) *TreeNode {
	var root *TreeNode
	for _, value := range values {
		root = InsertBST(root, value)
	}
	return root
}

// SearchBST 在二叉搜索树中查找值，未找到返回 nil。
func SearchBST(root *TreeNode, value int) *TreeNode {
	for root != nil {
		switch {
		case value == root.Val:
			return root
		case value < root.Val:
			root = root.Left
		default:
			root = root.Right
		}
	}
	return nil
}

// Inorder 返回中序遍历结果。
func Inorder(root *TreeNode) []int {
	values := make([]int, 0)
	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Left)
		values = append(values, node.Val)
		visit(node.Right)
	}
	visit(root)
	return values
}

// Preorder 返回前序遍历结果。
func Preorder(root *TreeNode) []int {
	values := make([]int, 0)
	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		visit(node.Left)
		visit(node.Right)
	}
	visit(root)
	return values
}

// Postorder 返回后序遍历结果。
func Postorder(root *TreeNode) []int {
	values := make([]int, 0)
	var visit func(node *TreeNode)
	visit = func(node *TreeNode) {
		if node == nil {
			return
		}
		visit(node.Left)
		visit(node.Right)
		values = append(values, node.Val)
	}
	visit(root)
	return values
}

// LevelOrder 返回层序遍历结果，每层一个切片。
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelSize:]
		ans = append(ans, level)
	}
	return ans
}

// MaxDepth 返回二叉树的最大深度。
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

// IsValidBST 判断是否为合法的二叉搜索树。
// 使用上下界传递，避免只比较父子节点导致的错误。
func IsValidBST(root *TreeNode) bool {
	var validate func(node *TreeNode, low, high int) bool
	validate = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}
		if node.Val <= low || node.Val >= high {
			return false
		}
		return validate(node.Left, low, node.Val) && validate(node.Right, node.Val, high)
	}
	return validate(root, math.MinInt, math.MaxInt)
}
