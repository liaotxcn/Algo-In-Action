// Package solution 提供常见排序算法的 Go 实现。
// 所有排序函数均为原地升序排序。
package solution

// BubbleSort 冒泡排序，时间复杂度 O(n²)，空间复杂度 O(1)。
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}

// SelectionSort 选择排序，时间复杂度 O(n²)，空间复杂度 O(1)。
func SelectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// InsertionSort 直接插入排序，时间复杂度 O(n²)，最好 O(n)。
func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

// ShellSort 希尔排序，时间复杂度依赖增量序列，最坏 O(n²)。
func ShellSort(nums []int) {
	n := len(nums)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			key := nums[i]
			j := i - gap
			for j >= 0 && nums[j] > key {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = key
		}
	}
}

// MergeSort 归并排序，时间复杂度 O(n log n)，空间复杂度 O(n)。
func MergeSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	tmp := make([]int, len(nums))
	mergeSortRange(nums, tmp, 0, len(nums)-1)
}

// mergeSortRange 递归排序闭区间 [left, right]。
func mergeSortRange(nums, tmp []int, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSortRange(nums, tmp, left, mid)
	mergeSortRange(nums, tmp, mid+1, right)

	i, j, write := left, mid+1, left
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[write] = nums[i]
			i++
		} else {
			tmp[write] = nums[j]
			j++
		}
		write++
	}
	for i <= mid {
		tmp[write] = nums[i]
		i++
		write++
	}
	for j <= right {
		tmp[write] = nums[j]
		j++
		write++
	}
	copy(nums[left:right+1], tmp[left:right+1])
}

// QuickSort 三路快速排序，平均 O(n log n)，最坏 O(n²)。
func QuickSort(nums []int) {
	quickSortRange(nums, 0, len(nums)-1)
}

func quickSortRange(nums []int, low, high int) {
	if low >= high {
		return
	}

	pivot := nums[low+(high-low)/2]
	lt, i, gt := low, low, high
	for i <= gt {
		switch {
		case nums[i] < pivot:
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		case nums[i] > pivot:
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		default:
			i++
		}
	}
	quickSortRange(nums, low, lt-1)
	quickSortRange(nums, gt+1, high)
}

// HeapSort 堆排序，时间复杂度 O(n log n)，空间复杂度 O(1)。
func HeapSort(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(nums, i, n)
	}
	for end := n - 1; end > 0; end-- {
		nums[0], nums[end] = nums[end], nums[0]
		siftDown(nums, 0, end)
	}
}

// siftDown 在 [0, size) 范围内维护大顶堆。
func siftDown(nums []int, index, size int) {
	for {
		largest := index
		left, right := 2*index+1, 2*index+2
		if left < size && nums[left] > nums[largest] {
			largest = left
		}
		if right < size && nums[right] > nums[largest] {
			largest = right
		}
		if largest == index {
			return
		}
		nums[index], nums[largest] = nums[largest], nums[index]
		index = largest
	}
}

// CountingSort 计数排序，适用于取值范围有限的整数。
// 时间复杂度 O(n+k)，空间复杂度 O(k)。
func CountingSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	minValue, maxValue := nums[0], nums[0]
	for _, value := range nums[1:] {
		minValue = min(minValue, value)
		maxValue = max(maxValue, value)
	}

	counts := make([]int, maxValue-minValue+1)
	for _, value := range nums {
		counts[value-minValue]++
	}

	write := 0
	for offset, count := range counts {
		for ; count > 0; count-- {
			nums[write] = offset + minValue
			write++
		}
	}
}

// BucketSort 桶排序，数据分布均匀时接近 O(n)。
func BucketSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	minValue, maxValue := nums[0], nums[0]
	for _, value := range nums[1:] {
		minValue = min(minValue, value)
		maxValue = max(maxValue, value)
	}
	if minValue == maxValue {
		return
	}

	bucketCount := max(1, n/5)
	buckets := make([][]int, bucketCount)
	span := float64(maxValue-minValue+1) / float64(bucketCount)
	for _, value := range nums {
		index := int(float64(value-minValue) / span)
		if index >= bucketCount {
			index = bucketCount - 1
		}
		buckets[index] = append(buckets[index], value)
	}

	write := 0
	for _, bucket := range buckets {
		InsertionSort(bucket)
		for _, value := range bucket {
			nums[write] = value
			write++
		}
	}
}

// RadixSort 基数排序（LSD），支持负数。
// 时间复杂度 O(d·(n+10))，空间复杂度 O(n)。
func RadixSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	minValue := nums[0]
	for _, value := range nums[1:] {
		minValue = min(minValue, value)
	}

	// 平移到非负区间，避免负数取模问题
	shifted := make([]uint64, n)
	var maxShifted uint64
	for i, value := range nums {
		current := uint64(int64(value) - int64(minValue))
		shifted[i] = current
		maxShifted = max(maxShifted, current)
	}

	buffer := make([]uint64, n)
	for exp := uint64(1); maxShifted/exp > 0; exp *= 10 {
		var counts [10]int
		for _, value := range shifted {
			counts[(value/exp)%10]++
		}
		for i := 1; i < len(counts); i++ {
			counts[i] += counts[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := (shifted[i] / exp) % 10
			counts[digit]--
			buffer[counts[digit]] = shifted[i]
		}
		copy(shifted, buffer)
	}

	for i, value := range shifted {
		nums[i] = int(int64(value) + int64(minValue))
	}
}
