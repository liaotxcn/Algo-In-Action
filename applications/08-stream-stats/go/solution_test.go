package solution

import (
	"reflect"
	"testing"
)

func TestCountMinSketch(t *testing.T) {
	sketch := NewCountMinSketch(256, 4)
	inputs := map[string]int{
		"apple":  10,
		"banana": 3,
		"cherry": 7,
	}
	for item, count := range inputs {
		sketch.Add([]byte(item), count)
	}

	for item, count := range inputs {
		if got := sketch.Estimate([]byte(item)); got < count {
			t.Fatalf("Estimate(%q) = %d, 不应小于真实值 %d", item, got, count)
		}
	}
	if got := sketch.Estimate([]byte("missing")); got != 0 {
		t.Fatalf("未出现元素估计值 = %d, want 0", got)
	}
}

func TestCountMinSketchMonotonic(t *testing.T) {
	sketch := NewCountMinSketch(64, 3)
	sketch.Add([]byte("key"), 2)
	before := sketch.Estimate([]byte("key"))
	sketch.Add([]byte("key"), 5)
	after := sketch.Estimate([]byte("key"))
	if after < before+5 {
		t.Fatalf("增量估计 = %d, want >= %d", after, before+5)
	}
}

func TestMovingAverage(t *testing.T) {
	average := NewMovingAverage(3)
	average.Add(1)
	average.Add(2)
	average.Add(3)
	if got := average.Average(); got != 2 {
		t.Fatalf("Average = %v, want 2", got)
	}

	average.Add(6) // 窗口变为 2, 3, 6
	if got := average.Average(); got != 11.0/3.0 {
		t.Fatalf("滚动后 Average = %v, want %v", got, 11.0/3.0)
	}
	if average.Len() != 3 {
		t.Fatalf("Len = %d, want 3", average.Len())
	}
}

func TestMovingAverageEmpty(t *testing.T) {
	average := NewMovingAverage(2)
	if average.Average() != 0 {
		t.Fatalf("空窗口平均值 = %v, want 0", average.Average())
	}
}

func TestReservoirSamplerDeterministic(t *testing.T) {
	first := NewReservoirSampler(5, 42)
	second := NewReservoirSampler(5, 42)
	for value := 0; value < 100; value++ {
		first.Add(value)
		second.Add(value)
	}

	if !reflect.DeepEqual(first.Sample(), second.Sample()) {
		t.Fatalf("相同种子应产生相同样本: %v vs %v", first.Sample(), second.Sample())
	}
	if first.Seen() != 100 {
		t.Fatalf("Seen = %d, want 100", first.Seen())
	}
}

func TestReservoirSamplerCapacity(t *testing.T) {
	sampler := NewReservoirSampler(3, 1)
	sampler.Add(10)
	sampler.Add(20)
	sampler.Add(30)
	if got := sampler.Sample(); !reflect.DeepEqual(got, []int{10, 20, 30}) {
		t.Fatalf("未超过容量时样本 = %v, want [10 20 30]", got)
	}

	for value := 40; value < 100; value++ {
		sampler.Add(value)
	}
	sample := sampler.Sample()
	if len(sample) != 3 {
		t.Fatalf("样本大小 = %d, want 3", len(sample))
	}
	for _, value := range sample {
		if value < 10 || value >= 100 {
			t.Fatalf("样本值 %d 不在输入范围内", value)
		}
	}
}
