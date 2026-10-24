// Package solution 提供哈希表相关算法的 Go 实现。
package solution

import (
	"slices"
	"strings"
)

// TwoSum 返回和为目标值的两个下标；不存在时返回 nil。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for index, value := range nums {
		if other, ok := seen[target-value]; ok {
			return []int{other, index}
		}
		seen[value] = index
	}
	return nil
}

// LongestConsecutive 返回最长连续序列的长度，要求 O(n)。
func LongestConsecutive(nums []int) int {
	values := make(map[int]struct{}, len(nums))
	for _, value := range nums {
		values[value] = struct{}{}
	}

	ans := 0
	for value := range values {
		// 只从序列起点开始扩展
		if _, hasPrev := values[value-1]; hasPrev {
			continue
		}
		length := 1
		for {
			if _, ok := values[value+length]; !ok {
				break
			}
			length++
		}
		ans = max(ans, length)
	}
	return ans
}

// GroupAnagrams 将字母异位词分组。
// 返回的分组与分组内元素均按字典序排列，便于稳定测试。
// 时间复杂度 O(n·k log k)，空间复杂度 O(n·k)。
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range strs {
		letters := []byte(word)
		slices.Sort(letters)
		key := string(letters)
		groups[key] = append(groups[key], word)
	}

	ans := make([][]string, 0, len(groups))
	for _, group := range groups {
		slices.Sort(group)
		ans = append(ans, group)
	}
	// 按分组内首个单词排序，保证跨语言输出稳定
	slices.SortFunc(ans, func(a, b []string) int { return strings.Compare(a[0], b[0]) })
	return ans
}

// SubarraySum 返回和为 k 的连续子数组个数。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func SubarraySum(nums []int, k int) int {
	prefixCount := map[int]int{0: 1}
	prefix, ans := 0, 0
	for _, value := range nums {
		prefix += value
		ans += prefixCount[prefix-k]
		prefixCount[prefix]++
	}
	return ans
}

// IsAnagram 判断两个字符串是否互为字母异位词。
// 使用长度 26 的计数数组，时间复杂度 O(n)，空间复杂度 O(1)。
func IsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(a); i++ {
		counts[a[i]-'a']++
		counts[b[i]-'a']--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

// NormalizeWord 返回按字母排序后的字符串，便于构建分组键。
// 该函数用于演示 strings 包的常见用法，不参与核心算法。
func NormalizeWord(word string) string {
	letters := strings.Split(word, "")
	slices.Sort(letters)
	return strings.Join(letters, "")
}
