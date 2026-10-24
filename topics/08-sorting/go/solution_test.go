package solution

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type sortCase struct {
	Name     string `json:"name"`
	Input    []int  `json:"input"`
	Expected []int  `json:"expected"`
}

type sortCases struct {
	Cases []sortCase `json:"cases"`
}

// loadSharedCases 读取 Go 与 Python 共用的排序用例。
func loadSharedCases(t *testing.T) []sortCase {
	t.Helper()
	data, err := os.ReadFile("../../../shared/testdata/sort_cases.json")
	if err != nil {
		t.Fatalf("读取共享排序用例失败: %v", err)
	}
	var payload sortCases
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("解析共享排序用例失败: %v", err)
	}
	return payload.Cases
}

func TestAllSorts(t *testing.T) {
	sorts := []struct {
		name string
		fn   func([]int)
	}{
		{"bubble", BubbleSort},
		{"selection", SelectionSort},
		{"insertion", InsertionSort},
		{"shell", ShellSort},
		{"merge", MergeSort},
		{"quick", QuickSort},
		{"heap", HeapSort},
		{"counting", CountingSort},
		{"bucket", BucketSort},
		{"radix", RadixSort},
	}

	for _, sort := range sorts {
		for _, tc := range loadSharedCases(t) {
			t.Run(sort.name+"/"+tc.Name, func(t *testing.T) {
				nums := make([]int, len(tc.Input))
				copy(nums, tc.Input)
				sort.fn(nums)
				if !reflect.DeepEqual(nums, tc.Expected) {
					t.Fatalf("%s(%v) = %v, want %v", sort.name, tc.Input, nums, tc.Expected)
				}
			})
		}
	}
}
