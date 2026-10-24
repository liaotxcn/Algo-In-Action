// Package solution 提供回溯算法经典问题的 Go 实现。
package solution

import (
	"slices"
	"strings"
)

// Permutations 返回不重复元素的全排列，顺序按输入顺序生成。
// 时间复杂度 O(n * n!)，空间复杂度 O(n)。
func Permutations(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			result = append(result, slices.Clone(path))
			return
		}
		for i, value := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, value)
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}

// PermutationsUnique 返回含重复元素的全排列（去重）。
// 先排序，再跳过同层重复元素，时间复杂度 O(n * n!)。
func PermutationsUnique(nums []int) [][]int {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)

	result := make([][]int, 0)
	path := make([]int, 0, len(sorted))
	used := make([]bool, len(sorted))
	var backtrack func()
	backtrack = func() {
		if len(path) == len(sorted) {
			result = append(result, slices.Clone(path))
			return
		}
		for i, value := range sorted {
			if used[i] {
				continue
			}
			if i > 0 && sorted[i] == sorted[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, value)
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}

// Combine 返回 1 到 n 中所有长度为 k 的组合，按字典序生成。
func Combine(n, k int) [][]int {
	if k <= 0 || k > n {
		return [][]int{}
	}

	result := make([][]int, 0)
	path := make([]int, 0, k)
	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			result = append(result, slices.Clone(path))
			return
		}
		// 剩余元素不足以填满时提前剪枝
		for value := start; value <= n-(k-len(path))+1; value++ {
			path = append(path, value)
			backtrack(value + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return result
}

// Subsets 返回数组的全部子集（幂集）。
func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0, len(nums))
	var backtrack func(start int)
	backtrack = func(start int) {
		result = append(result, slices.Clone(path))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}

// SolveNQueens 返回 n 皇后问题的所有解，'.' 表示空位，'Q' 表示皇后。
func SolveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}

	result := make([][]string, 0)
	cols := make([]bool, n)
	diag := make([]bool, 2*n-1)
	antiDiag := make([]bool, 2*n-1)
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			result = append(result, renderBoard(queens))
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag[row+col] || antiDiag[row-col+n-1] {
				continue
			}
			cols[col], diag[row+col], antiDiag[row-col+n-1] = true, true, true
			queens[row] = col
			backtrack(row + 1)
			queens[row] = -1
			cols[col], diag[row+col], antiDiag[row-col+n-1] = false, false, false
		}
	}
	backtrack(0)
	return result
}

// renderBoard 将列下标数组渲染为棋盘字符串。
func renderBoard(queens []int) []string {
	board := make([]string, len(queens))
	for row, col := range queens {
		var builder strings.Builder
		builder.WriteString(strings.Repeat(".", col))
		builder.WriteString("Q")
		builder.WriteString(strings.Repeat(".", len(queens)-col-1))
		board[row] = builder.String()
	}
	return board
}
