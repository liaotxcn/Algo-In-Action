// Package solution 提供有序数组上的搜索算法实现。
package solution

// BinarySearch 在有序数组中查找 target，返回任意一个匹配下标；未找到返回 -1。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}

// LowerBound 返回第一个大于等于 target 的下标；不存在时返回 len(nums)。
// 时间复杂度 O(log n)，空间复杂度 O(1)。
func LowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// UpperBound 返回第一个大于 target 的下标；不存在时返回 len(nums)。
func UpperBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// SearchRange 返回 target 在有序数组中的起止下标；不存在返回 [-1, -1]。
func SearchRange(nums []int, target int) [2]int {
	start := LowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return [2]int{-1, -1}
	}
	end := UpperBound(nums, target) - 1
	return [2]int{start, end}
}

// InterpolationSearch 在分布均匀的有序数组中插值查找。
// 平均时间复杂度 O(log log n)，最坏 O(n)。
func InterpolationSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right && target >= nums[left] && target <= nums[right] {
		if nums[left] == nums[right] {
			if nums[left] == target {
				return left
			}
			return -1
		}

		// 按数值比例估计目标位置
		pos := left + (target-nums[left])*(right-left)/(nums[right]-nums[left])
		switch {
		case nums[pos] == target:
			return pos
		case nums[pos] < target:
			left = pos + 1
		default:
			right = pos - 1
		}
	}
	return -1
}
