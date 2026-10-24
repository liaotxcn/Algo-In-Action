// Package solution 提供图相关算法的 Go 实现。
package solution

import (
	"container/heap"
	"math"
)

// Edge 表示一条有向边。
type Edge struct {
	To     int
	Weight int
}

// Graph 使用邻接表存储图。
type Graph struct {
	Nodes int
	Adj   [][]Edge
}

// NewGraph 创建包含 nodes 个节点、初始无边界的图。
func NewGraph(nodes int) *Graph {
	return &Graph{Nodes: nodes, Adj: make([][]Edge, nodes)}
}

// AddEdge 添加一条有向边。
func (g *Graph) AddEdge(from, to, weight int) {
	if from < 0 || from >= g.Nodes || to < 0 || to >= g.Nodes {
		return
	}
	g.Adj[from] = append(g.Adj[from], Edge{To: to, Weight: weight})
}

// AddUndirectedEdge 添加一条无向边。
func (g *Graph) AddUndirectedEdge(a, b, weight int) {
	g.AddEdge(a, b, weight)
	g.AddEdge(b, a, weight)
}

// BFS 返回从 start 出发的广度优先遍历顺序。
func (g *Graph) BFS(start int) []int {
	if start < 0 || start >= g.Nodes {
		return []int{}
	}

	visited := make([]bool, g.Nodes)
	visited[start] = true
	queue := []int{start}
	order := make([]int, 0, g.Nodes)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, edge := range g.Adj[node] {
			if !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}
	return order
}

// DFS 返回从 start 出发的深度优先遍历顺序（递归实现）。
func (g *Graph) DFS(start int) []int {
	if start < 0 || start >= g.Nodes {
		return []int{}
	}

	visited := make([]bool, g.Nodes)
	order := make([]int, 0, g.Nodes)
	var visit func(node int)
	visit = func(node int) {
		visited[node] = true
		order = append(order, node)
		for _, edge := range g.Adj[node] {
			if !visited[edge.To] {
				visit(edge.To)
			}
		}
	}
	visit(start)
	return order
}

// Dijkstra 计算从 start 到所有节点的最短距离与前驱。
// 边权必须非负；不可达节点的距离为 math.MaxInt，前驱为 -1。
func (g *Graph) Dijkstra(start int) ([]int, []int) {
	dist := make([]int, g.Nodes)
	prev := make([]int, g.Nodes)
	for i := range dist {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	if start < 0 || start >= g.Nodes {
		return dist, prev
	}
	dist[start] = 0

	pq := &distanceQueue{{node: start, distance: 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		item := heap.Pop(pq).(distanceItem)
		if item.distance > dist[item.node] {
			continue // 过期记录
		}
		for _, edge := range g.Adj[item.node] {
			next := item.distance + edge.Weight
			if next < dist[edge.To] {
				dist[edge.To] = next
				prev[edge.To] = item.node
				heap.Push(pq, distanceItem{node: edge.To, distance: next})
			}
		}
	}
	return dist, prev
}

// BuildPath 根据前驱数组重建从起点到 target 的路径；不可达返回空切片。
func (g *Graph) BuildPath(prev []int, target int) []int {
	if target < 0 || target >= len(prev) {
		return []int{}
	}

	path := make([]int, 0)
	for node := target; node != -1; node = prev[node] {
		path = append(path, node)
	}
	// 反转
	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}
	return path
}

// TopologicalSort 使用 Kahn 算法返回拓扑序；存在环时返回 false。
func (g *Graph) TopologicalSort() ([]int, bool) {
	indegree := make([]int, g.Nodes)
	for from := 0; from < g.Nodes; from++ {
		for _, edge := range g.Adj[from] {
			indegree[edge.To]++
		}
	}

	queue := make([]int, 0, g.Nodes)
	for node, degree := range indegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	order := make([]int, 0, g.Nodes)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, edge := range g.Adj[node] {
			indegree[edge.To]--
			if indegree[edge.To] == 0 {
				queue = append(queue, edge.To)
			}
		}
	}

	if len(order) != g.Nodes {
		return []int{}, false
	}
	return order, true
}

// distanceItem 是 Dijkstra 优先队列元素。
type distanceItem struct {
	node     int
	distance int
}

// distanceQueue 实现 container/heap.Interface（小顶堆）。
type distanceQueue []distanceItem

func (q distanceQueue) Len() int           { return len(q) }
func (q distanceQueue) Less(i, j int) bool { return q[i].distance < q[j].distance }
func (q distanceQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *distanceQueue) Push(x any)        { *q = append(*q, x.(distanceItem)) }
func (q *distanceQueue) Pop() any {
	old := *q
	last := len(old) - 1
	item := old[last]
	*q = old[:last]
	return item
}
