// Package solution 提供贪心算法经典问题的 Go 实现。
package solution

import "slices"

// minEndTime 作为哨兵值，保证第一个活动一定能被选中。
const minEndTime = -1 << 63

// ActivitySelection 返回最多可参加的活动下标，按结束时间升序选择。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func ActivitySelection(start, end []int) []int {
	n := len(start)
	if n == 0 || n != len(end) {
		return []int{}
	}

	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	slices.SortFunc(order, func(a, b int) int { return end[a] - end[b] })

	selected := make([]int, 0, n)
	lastEnd := minEndTime
	for _, index := range order {
		if start[index] >= lastEnd {
			selected = append(selected, index)
			lastEnd = end[index]
		}
	}
	return selected
}

type item struct {
	weight int
	value  int
	ratio  float64
}

// FractionalKnapsack 返回分数背包能装入的最大价值。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func FractionalKnapsack(weights, values []int, capacity float64) float64 {
	if capacity <= 0 || len(weights) == 0 || len(weights) != len(values) {
		return 0
	}

	items := make([]item, len(weights))
	for i, weight := range weights {
		ratio := 0.0
		if weight > 0 {
			ratio = float64(values[i]) / float64(weight)
		}
		items[i] = item{weight: weight, value: values[i], ratio: ratio}
	}
	slices.SortFunc(items, func(a, b item) int {
		switch {
		case a.ratio > b.ratio:
			return -1
		case a.ratio < b.ratio:
			return 1
		default:
			return 0
		}
	})

	total := 0.0
	remaining := capacity
	for _, current := range items {
		if remaining <= 0 {
			break
		}
		take := min(float64(current.weight), remaining)
		total += take * current.ratio
		remaining -= take
	}
	return total
}

// JumpGame 判断能否从下标 0 跳到最后一个下标。
// 维护当前可达的最远位置，时间复杂度 O(n)。
func JumpGame(nums []int) bool {
	reach := 0
	for i, jump := range nums {
		if i > reach {
			return false
		}
		reach = max(reach, i+jump)
	}
	return true
}

// MinJumps 返回跳到最后一个下标的最少步数；不可达返回 -1。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func MinJumps(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumps, currentEnd, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currentEnd {
			jumps++
			currentEnd = farthest
			if currentEnd >= len(nums)-1 {
				return jumps
			}
		}
	}
	return -1
}

// FindContentChildren 返回最多能满足的孩子数量（贪心分配饼干）。
func FindContentChildren(greed, cookies []int) int {
	slices.Sort(greed)
	slices.Sort(cookies)

	child, cookie := 0, 0
	for child < len(greed) && cookie < len(cookies) {
		if cookies[cookie] >= greed[child] {
			child++
		}
		cookie++
	}
	return child
}
