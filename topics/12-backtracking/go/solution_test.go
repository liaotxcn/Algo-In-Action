package solution

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	got := Permutations([]int{1, 2, 3})
	want := [][]int{
		{1, 2, 3}, {1, 3, 2},
		{2, 1, 3}, {2, 3, 1},
		{3, 1, 2}, {3, 2, 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Permutations = %v, want %v", got, want)
	}
	if got := Permutations([]int{1}); len(got) != 1 {
		t.Fatalf("单元素排列数量 = %d, want 1", len(got))
	}
}

func TestPermutationsUnique(t *testing.T) {
	got := PermutationsUnique([]int{1, 1, 2})
	want := [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("PermutationsUnique = %v, want %v", got, want)
	}
}

func TestCombine(t *testing.T) {
	got := Combine(4, 2)
	want := [][]int{
		{1, 2}, {1, 3}, {1, 4},
		{2, 3}, {2, 4}, {3, 4},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Combine = %v, want %v", got, want)
	}
	if got := Combine(4, 0); len(got) != 0 {
		t.Fatalf("k=0 应返回空结果, got %v", got)
	}
}

func TestSubsets(t *testing.T) {
	got := Subsets([]int{1, 2})
	want := [][]int{{}, {1}, {1, 2}, {2}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Subsets = %v, want %v", got, want)
	}
}

func TestSolveNQueens(t *testing.T) {
	if got := SolveNQueens(4); len(got) != 2 {
		t.Fatalf("4 皇后解数量 = %d, want 2", len(got))
	}
	want := []string{".Q..", "...Q", "Q...", "..Q."}
	if got := SolveNQueens(4); !reflect.DeepEqual(got[0], want) {
		t.Fatalf("首个解 = %v, want %v", got[0], want)
	}
	if got := SolveNQueens(0); len(got) != 0 {
		t.Fatalf("n=0 应返回空结果, got %v", got)
	}
}
