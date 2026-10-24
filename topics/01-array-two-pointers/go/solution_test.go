package solution

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

type maxAreaCase struct {
	Name     string `json:"name"`
	Height   []int  `json:"height"`
	Expected int    `json:"expected"`
}

type rotateCase struct {
	Name     string `json:"name"`
	Nums     []int  `json:"nums"`
	K        int    `json:"k"`
	Expected []int  `json:"expected"`
}

type sharedCases struct {
	MaxArea []maxAreaCase `json:"max_area"`
	Rotate  []rotateCase  `json:"rotate"`
}

// loadCases 读取与 Python 测试共享的用例文件。
func loadCases(t *testing.T) sharedCases {
	t.Helper()
	data, err := os.ReadFile("../cases/cases.json")
	if err != nil {
		t.Fatalf("读取共享用例失败: %v", err)
	}
	var cases sharedCases
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("解析共享用例失败: %v", err)
	}
	return cases
}

func TestRotateArraySharedCases(t *testing.T) {
	for _, tc := range loadCases(t).Rotate {
		t.Run(tc.Name, func(t *testing.T) {
			// 使用非 nil 切片，便于与 JSON 中的空数组比较
			nums := make([]int, len(tc.Nums))
			copy(nums, tc.Nums)
			RotateArray(nums, tc.K)
			if !reflect.DeepEqual(nums, tc.Expected) {
				t.Fatalf("RotateArray(%v, %d) = %v, want %v", tc.Nums, tc.K, nums, tc.Expected)
			}
		})
	}
}

func TestMaxAreaSharedCases(t *testing.T) {
	for _, tc := range loadCases(t).MaxArea {
		t.Run(tc.Name, func(t *testing.T) {
			if got := MaxArea(tc.Height); got != tc.Expected {
				t.Fatalf("MaxArea(%v) = %d, want %d", tc.Height, got, tc.Expected)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"classic", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"all_zero", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"no_solution", []int{1, 2, 3}, [][]int{}},
		{"too_short", []int{0, 1}, [][]int{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			if got := ThreeSum(nums); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("ThreeSum(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	want := []int{1, 3, 12, 0, 0}
	if !reflect.DeepEqual(nums, want) {
		t.Fatalf("MoveZeroes = %v, want %v", nums, want)
	}
}

func TestTrap(t *testing.T) {
	cases := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{}, 0},
	}
	for _, tc := range cases {
		if got := Trap(tc.height); got != tc.want {
			t.Fatalf("Trap(%v) = %d, want %d", tc.height, got, tc.want)
		}
	}
}

func TestMergeIntervals(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	want := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if got := MergeIntervals(input); !reflect.DeepEqual(got, want) {
		t.Fatalf("MergeIntervals = %v, want %v", got, want)
	}
	// 校验不修改入参
	if !reflect.DeepEqual(input, [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}) {
		t.Fatalf("MergeIntervals 不应修改入参, got %v", input)
	}
}

func TestProductExceptSelf(t *testing.T) {
	want := []int{24, 12, 8, 6}
	if got := ProductExceptSelf([]int{1, 2, 3, 4}); !reflect.DeepEqual(got, want) {
		t.Fatalf("ProductExceptSelf = %v, want %v", got, want)
	}
}

func TestMaxSubArray(t *testing.T) {
	if got := MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); got != 6 {
		t.Fatalf("MaxSubArray = %d, want 6", got)
	}
	if got := MaxSubArray([]int{-3, -1, -2}); got != -1 {
		t.Fatalf("MaxSubArray 全负数 = %d, want -1", got)
	}
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if got := SpiralOrder(matrix); !reflect.DeepEqual(got, want) {
		t.Fatalf("SpiralOrder = %v, want %v", got, want)
	}
	if got := SpiralOrder([][]int{}); len(got) != 0 {
		t.Fatalf("空矩阵应返回空切片, got %v", got)
	}
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	want := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	SetZeroes(matrix)
	if !reflect.DeepEqual(matrix, want) {
		t.Fatalf("SetZeroes = %v, want %v", matrix, want)
	}
}

func TestRotateMatrix(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	RotateMatrix(matrix)
	if !reflect.DeepEqual(matrix, want) {
		t.Fatalf("RotateMatrix = %v, want %v", matrix, want)
	}
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	if !SearchMatrix(matrix, 5) {
		t.Fatal("SearchMatrix(5) 应返回 true")
	}
	if SearchMatrix(matrix, 20) {
		t.Fatal("SearchMatrix(20) 应返回 false")
	}
}
