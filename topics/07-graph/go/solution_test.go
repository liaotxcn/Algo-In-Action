package solution

import (
	"math"
	"reflect"
	"testing"
)

// buildWeightedGraph 构建用于最短路径测试的无向图。
func buildWeightedGraph() *Graph {
	g := NewGraph(6)
	g.AddUndirectedEdge(0, 1, 4)
	g.AddUndirectedEdge(0, 2, 2)
	g.AddUndirectedEdge(1, 2, 1)
	g.AddUndirectedEdge(1, 3, 5)
	g.AddUndirectedEdge(2, 3, 8)
	g.AddUndirectedEdge(2, 4, 10)
	g.AddUndirectedEdge(3, 4, 2)
	g.AddUndirectedEdge(3, 5, 6)
	g.AddUndirectedEdge(4, 5, 3)
	return g
}

func TestBFSAndDFS(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)

	if got := g.BFS(0); !reflect.DeepEqual(got, []int{0, 1, 2, 3}) {
		t.Fatalf("BFS = %v, want [0 1 2 3]", got)
	}
	if got := g.DFS(0); !reflect.DeepEqual(got, []int{0, 1, 3, 2}) {
		t.Fatalf("DFS = %v, want [0 1 3 2]", got)
	}
}

func TestDijkstraAndPath(t *testing.T) {
	g := buildWeightedGraph()
	dist, prev := g.Dijkstra(0)
	want := []int{0, 3, 2, 8, 10, 13}
	if !reflect.DeepEqual(dist, want) {
		t.Fatalf("dist = %v, want %v", dist, want)
	}
	// 最短路径为 0 -> 2 -> 1 -> 3 -> 4 -> 5，总权重 2+1+5+2+3=13
	if got := g.BuildPath(prev, 5); !reflect.DeepEqual(got, []int{0, 2, 1, 3, 4, 5}) {
		t.Fatalf("path = %v, want [0 2 1 3 4 5]", got)
	}
}

func TestDijkstraUnreachable(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 5)
	dist, prev := g.Dijkstra(0)
	if dist[2] != math.MaxInt {
		t.Fatalf("不可达节点距离 = %d, want MaxInt", dist[2])
	}
	if got := g.BuildPath(prev, 2); len(got) != 1 || got[0] != 2 {
		t.Fatalf("不可达路径 = %v, want [2]", got)
	}
}

func TestTopologicalSort(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	order, ok := g.TopologicalSort()
	if !ok {
		t.Fatal("DAG 应能完成拓扑排序")
	}
	if !reflect.DeepEqual(order, []int{0, 1, 2, 3, 4}) {
		t.Fatalf("拓扑序 = %v, want [0 1 2 3 4]", order)
	}

	cyclic := NewGraph(2)
	cyclic.AddEdge(0, 1, 1)
	cyclic.AddEdge(1, 0, 1)
	if _, ok := cyclic.TopologicalSort(); ok {
		t.Fatal("有环图应返回 false")
	}
}
