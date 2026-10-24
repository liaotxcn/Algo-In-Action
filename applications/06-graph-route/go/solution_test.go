package solution

import (
	"reflect"
	"testing"
)

// buildRouter 构建用于路线规划的测试图。
func buildRouter() *Router {
	router := NewRouter(5)
	router.AddRoute(0, 1, 10)
	router.AddRoute(0, 2, 1)
	router.AddRoute(2, 1, 1)
	router.AddRoute(1, 3, 2)
	router.AddRoute(2, 4, 20)
	router.AddRoute(4, 3, 1)
	return router
}

func TestShortestPath(t *testing.T) {
	router := buildRouter()

	path, distance, ok := router.ShortestPath(0, 3)
	if !ok {
		t.Fatal("0 -> 3 应可达")
	}
	if want := []int{0, 2, 1, 3}; !reflect.DeepEqual(path, want) {
		t.Fatalf("最短路径 = %v, want %v", path, want)
	}
	if distance != 4 {
		t.Fatalf("最短距离 = %d, want 4", distance)
	}
}

func TestFewestHops(t *testing.T) {
	router := buildRouter()

	path, ok := router.FewestHops(0, 3)
	if !ok {
		t.Fatal("0 -> 3 应可达")
	}
	if want := []int{0, 1, 3}; !reflect.DeepEqual(path, want) {
		t.Fatalf("最少跳数路径 = %v, want %v", path, want)
	}
}

func TestUnreachable(t *testing.T) {
	router := buildRouter()
	if _, _, ok := router.ShortestPath(3, 0); ok {
		t.Fatal("3 -> 0 不可达")
	}
	if _, ok := router.FewestHops(3, 0); ok {
		t.Fatal("3 -> 0 不可达")
	}
	if router.Reachable(0, 3) == false {
		t.Fatal("0 -> 3 应可达")
	}
}

func TestInvalidNodes(t *testing.T) {
	router := buildRouter()
	if _, _, ok := router.ShortestPath(-1, 3); ok {
		t.Fatal("非法起点应返回 false")
	}
	if _, ok := router.FewestHops(0, 99); ok {
		t.Fatal("非法终点应返回 false")
	}
}
