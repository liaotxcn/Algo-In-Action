package solution

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	if !reflect.DeepEqual(got, []int{0, 1}) {
		t.Fatalf("TwoSum = %v, want [0 1]", got)
	}
	if got := TwoSum([]int{3, 3}, 6); !reflect.DeepEqual(got, []int{0, 1}) {
		t.Fatalf("TwoSum = %v, want [0 1]", got)
	}
	if got := TwoSum([]int{1, 2}, 10); got != nil {
		t.Fatalf("无解时应返回 nil, got %v", got)
	}
}

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
		{[]int{1, 1, 1}, 1},
	}
	for _, tc := range cases {
		if got := LongestConsecutive(tc.nums); got != tc.want {
			t.Fatalf("LongestConsecutive(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GroupAnagrams = %v, want %v", got, want)
	}
	if got := GroupAnagrams(nil); len(got) != 0 {
		t.Fatalf("空输入应返回空结果, got %v", got)
	}
}

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, -1, 0}, 0, 3},
	}
	for _, tc := range cases {
		if got := SubarraySum(tc.nums, tc.k); got != tc.want {
			t.Fatalf("SubarraySum(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	if !IsAnagram("anagram", "nagaram") {
		t.Fatal("anagram 与 nagaram 应互为异位词")
	}
	if IsAnagram("rat", "car") {
		t.Fatal("rat 与 car 不应互为异位词")
	}
}

func TestNormalizeWord(t *testing.T) {
	if got := NormalizeWord("eat"); got != "aet" {
		t.Fatalf("NormalizeWord = %q, want %q", got, "aet")
	}
}
