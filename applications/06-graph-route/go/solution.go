// Package solution 提供路线规划（最短距离与最少跳数）的 Go 实现。
package solution

import (
	"container/heap"
	"math"
)

// routeEdge 表示一条有向路线。
type routeEdge struct {
	to       int
	distance int
}

// Router 使用邻接表存储路线图。
type Router struct {
	nodes int
	adj   [][]routeEdge
}

// NewRouter 创建包含 nodes 个地点的路线图。
func NewRouter(nodes int) *Router {
	if nodes < 0 {
		nodes = 0
	}
	return &Router{nodes: nodes, adj: make([][]routeEdge, nodes)}
}

// AddRoute 添加一条有向路线；节点越界时忽略。
func (r *Router) AddRoute(from, to, distance int) {
	if from < 0 || from >= r.nodes || to < 0 || to >= r.nodes {
		return
	}
	r.adj[from] = append(r.adj[from], routeEdge{to: to, distance: distance})
}

// ShortestPath 使用 Dijkstra 返回距离最短的路径与总距离。
// 不可达时返回 nil、0、false。
func (r *Router) ShortestPath(from, to int) ([]int, int, bool) {
	if !r.valid(from) || !r.valid(to) {
		return nil, 0, false
	}

	dist := make([]int, r.nodes)
	prev := make([]int, r.nodes)
	for i := range dist {
		dist[i] = math.MaxInt
		prev[i] = -1
	}
	dist[from] = 0

	queue := &routeQueue{{node: from, distance: 0}}
	heap.Init(queue)
	for queue.Len() > 0 {
		current := heap.Pop(queue).(routeItem)
		if current.distance > dist[current.node] {
			continue
		}
		if current.node == to {
			break
		}
		for _, edge := range r.adj[current.node] {
			next := current.distance + edge.distance
			if next < dist[edge.to] {
				dist[edge.to] = next
				prev[edge.to] = current.node
				heap.Push(queue, routeItem{node: edge.to, distance: next})
			}
		}
	}

	if dist[to] == math.MaxInt {
		return nil, 0, false
	}
	return rebuildPath(prev, from, to), dist[to], true
}

// FewestHops 使用 BFS 返回经过地点最少的路径。
// 不可达时返回 nil、false。
func (r *Router) FewestHops(from, to int) ([]int, bool) {
	if !r.valid(from) || !r.valid(to) {
		return nil, false
	}

	prev := make([]int, r.nodes)
	for i := range prev {
		prev[i] = -1
	}
	visited := make([]bool, r.nodes)
	visited[from] = true
	queue := []int{from}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == to {
			return rebuildPath(prev, from, to), true
		}
		for _, edge := range r.adj[node] {
			if visited[edge.to] {
				continue
			}
			visited[edge.to] = true
			prev[edge.to] = node
			queue = append(queue, edge.to)
		}
	}
	return nil, false
}

// Reachable 判断从 from 是否能到达 to。
func (r *Router) Reachable(from, to int) bool {
	_, ok := r.FewestHops(from, to)
	return ok
}

// valid 判断节点下标是否合法。
func (r *Router) valid(node int) bool {
	return node >= 0 && node < r.nodes
}

// rebuildPath 根据前驱数组重建路径。
func rebuildPath(prev []int, from, to int) []int {
	path := make([]int, 0)
	for node := to; node != -1; node = prev[node] {
		path = append(path, node)
		if node == from {
			break
		}
	}
	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}
	return path
}

// routeItem 是 Dijkstra 优先队列元素。
type routeItem struct {
	node     int
	distance int
}

// routeQueue 实现小顶堆，距离相同时按节点编号排序以保证结果确定。
type routeQueue []routeItem

func (q routeQueue) Len() int { return len(q) }
func (q routeQueue) Less(i, j int) bool {
	if q[i].distance != q[j].distance {
		return q[i].distance < q[j].distance
	}
	return q[i].node < q[j].node
}
func (q routeQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *routeQueue) Push(x any)   { *q = append(*q, x.(routeItem)) }
func (q *routeQueue) Pop() any {
	old := *q
	last := len(old) - 1
	value := old[last]
	*q = old[:last]
	return value
}
