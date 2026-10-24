// Package solution 提供数组、双指针与矩阵相关算法的 Go 实现。
package solution

import "slices"

// RotateArray 将切片原地向右轮转 k 步。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func RotateArray(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	// 处理 k 为负数或大于 n 的情况
	k = ((k % n) + n) % n
	if k == 0 {
		return
	}

	// 整体反转后，分别反转前 k 个和后 n-k 个元素
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

// MaxArea 返回两条竖线与 x 轴围成的最大容器面积。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func MaxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		ans = max(ans, area)

		// 移动较矮的一侧，才有可能获得更大的面积
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

// ThreeSum 返回所有和为 0 且不重复的三元组。
// 会原地排序传入的切片。时间复杂度 O(n²)，空间复杂度 O(1)（不计结果）。
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	slices.Sort(nums)

	ans := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		// 跳过重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 当前最小值已经使三数之和大于 0，后续不可能更小
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 当前最大值仍然小于 0，当前数不可能构成解
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum > 0:
				right--
			case sum < 0:
				left++
			default:
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的第二个数
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳过重复的第三个数
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return ans
}

// MoveZeroes 将所有 0 原地移动到切片末尾，保持非零元素的相对顺序。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func MoveZeroes(nums []int) {
	write := 0
	for _, value := range nums {
		if value != 0 {
			nums[write] = value
			write++
		}
	}
	clear(nums[write:])
}

// Trap 计算柱状图能接住的雨水总量。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func Trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	preMax, sufMax := 0, 0

	for left < right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])

		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			ans += sufMax - height[right]
			right--
		}
	}
	return ans
}

// MergeIntervals 合并所有重叠区间，返回按起点升序排列的新切片。
// 不修改传入的区间。时间复杂度 O(n log n)，空间复杂度 O(n)。
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)
	slices.SortFunc(sorted, func(a, b []int) int { return a[0] - b[0] })

	ans := make([][]int, 0, len(sorted))
	for _, interval := range sorted {
		last := len(ans) - 1
		if last >= 0 && interval[0] <= ans[last][1] {
			ans[last][1] = max(ans[last][1], interval[1])
			continue
		}
		ans = append(ans, []int{interval[0], interval[1]})
	}
	return ans
}

// ProductExceptSelf 返回除自身以外所有元素的乘积，不使用除法。
// 时间复杂度 O(n)，空间复杂度 O(n)（结果数组）。
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// suffix[i] 表示 nums[i+1:] 的乘积
	suffix := make([]int, n)
	suffix[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	prefix := 1
	for i, value := range nums {
		suffix[i] *= prefix
		prefix *= value
	}
	return suffix
}

// MaxSubArray 返回和最大的连续子数组之和（数组非空时）。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := nums[0]
	prefix, minPrefix := 0, 0
	for _, value := range nums {
		prefix += value
		ans = max(ans, prefix-minPrefix)
		minPrefix = min(minPrefix, prefix)
	}
	return ans
}

// SpiralOrder 按顺时针螺旋顺序返回矩阵中的所有元素。
// 时间复杂度 O(mn)，空间复杂度 O(mn)（结果数组）。
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	ans := make([]int, 0, len(matrix)*len(matrix[0]))

	for top <= bottom && left <= right {
		for col := left; col <= right; col++ {
			ans = append(ans, matrix[top][col])
		}
		top++

		for row := top; row <= bottom; row++ {
			ans = append(ans, matrix[row][right])
		}
		right--

		if top <= bottom {
			for col := right; col >= left; col-- {
				ans = append(ans, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				ans = append(ans, matrix[row][left])
			}
			left++
		}
	}
	return ans
}

// SetZeroes 将矩阵中 0 所在的行和列全部置零，使用 O(1) 额外空间。
// 时间复杂度 O(mn)，空间复杂度 O(1)。
func SetZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	rows, cols := len(matrix), len(matrix[0])
	firstRowHasZero := slices.Contains(matrix[0], 0)

	// 用第一行和第一列记录需要置零的行列
	for row := 1; row < rows; row++ {
		for col, value := range matrix[row] {
			if value == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// 根据标记置零；从右向左处理，避免第一列标记被提前覆盖
	for row := 1; row < rows; row++ {
		for col := cols - 1; col >= 0; col-- {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
	}

	if firstRowHasZero {
		clear(matrix[0])
	}
}

// RotateMatrix 将 n×n 矩阵顺时针旋转 90 度，原地完成。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func RotateMatrix(matrix [][]int) {
	n := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < n; j++ {
			row[j], matrix[j][i] = matrix[j][i], row[j]
		}
		slices.Reverse(row)
	}
}

// SearchMatrix 在每行每列均升序的矩阵中查找 target。
// 时间复杂度 O(m+n)，空间复杂度 O(1)。
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		switch {
		case matrix[row][col] == target:
			return true
		case matrix[row][col] < target:
			row++
		default:
			col--
		}
	}
	return false
}
