package solution

import (
	"reflect"
	"testing"
)

func TestInsertAndSearchBST(t *testing.T) {
	root := BuildBST([]int{5, 3, 8, 1, 4, 7, 9, 3})
	if got := Inorder(root); !reflect.DeepEqual(got, []int{1, 3, 4, 5, 7, 8, 9}) {
		t.Fatalf("中序遍历 = %v", got)
	}
	if node := SearchBST(root, 4); node == nil || node.Val != 4 {
		t.Fatalf("SearchBST(4) = %v, want 4", node)
	}
	if node := SearchBST(root, 6); node != nil {
		t.Fatalf("SearchBST(6) = %v, want nil", node)
	}
}

func TestTraversals(t *testing.T) {
	root := BuildBST([]int{4, 2, 6, 1, 3, 5, 7})
	cases := []struct {
		name string
		got  []int
		want []int
	}{
		{"preorder", Preorder(root), []int{4, 2, 1, 3, 6, 5, 7}},
		{"inorder", Inorder(root), []int{1, 2, 3, 4, 5, 6, 7}},
		{"postorder", Postorder(root), []int{1, 3, 2, 5, 7, 6, 4}},
	}
	for _, tc := range cases {
		if !reflect.DeepEqual(tc.got, tc.want) {
			t.Fatalf("%s = %v, want %v", tc.name, tc.got, tc.want)
		}
	}
}

func TestLevelOrder(t *testing.T) {
	// 手工构造题目中的树：3 -> 9,20 -> 15,7
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	want := [][]int{{3}, {9, 20}, {15, 7}}
	if got := LevelOrder(root); !reflect.DeepEqual(got, want) {
		t.Fatalf("LevelOrder = %v, want %v", got, want)
	}
	if got := LevelOrder(nil); len(got) != 0 {
		t.Fatalf("空树层序遍历应为空, got %v", got)
	}
}

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	if got := MaxDepth(root); got != 3 {
		t.Fatalf("MaxDepth = %d, want 3", got)
	}
	if got := MaxDepth(nil); got != 0 {
		t.Fatalf("空树深度 = %d, want 0", got)
	}
}

func TestIsValidBST(t *testing.T) {
	if !IsValidBST(BuildBST([]int{2, 1, 3})) {
		t.Fatal("合法 BST 应返回 true")
	}
	invalid := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	if IsValidBST(invalid) {
		t.Fatal("非法 BST 应返回 false")
	}
}
