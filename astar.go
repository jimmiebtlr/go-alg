package alg

import (
	"container/heap"
)

type AStar struct {
	// Where to start, where to end<
	start  *Node
	goal   *Node
	solved bool

	frontier *AStarNodeHeap
}

type AStarNodeHeap []*Node

func (n AStarNodeHeap) Len() int           { return len(n) }
func (n AStarNodeHeap) Less(i, j int) bool { return (n[i].PathCost +n[i].HieuristicCost) < (n[j].PathCost + n[j].HieuristicCost) }
func (n AStarNodeHeap) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (n *AStarNodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*n = append(*n, x.(*Node))
}

func (n *AStarNodeHeap) Pop() interface{} {
	old := *n
	length := len(old)
	x := old[length-1]
	*n = old[0 : length-1]
	return x
}

func NewAStar() (astar AStar) {
	astar.frontier = &AStarNodeHeap{}
	heap.Init(astar.frontier)
	astar.solved = false
	return astar
}

func (astar *AStar) SetStart(n *Node) {
	astar.start = n
	astar.start.PathCost = 0
	heap.Push(astar.frontier, n)
}

func (astar *AStar) SetGoal(n *Node) {
	astar.goal = n
}

func (astar *AStar) expandFrontNode() (goalReached bool) {
	frontNode := astar.frontier.Pop().(*Node)

	if astar.goal == frontNode {
		astar.solved = true
		return true
	}

	for _, adj := range frontNode.Adj() {
		if adj.PathCost > (frontNode.PathCost + frontNode.TraversalCost(adj)) {
			adj.Prev = frontNode
			adj.PathCost = frontNode.PathCost + frontNode.TraversalCost(adj)
			astar.frontier.Push(adj)
		}
	}
	return false
}

func (astar *AStar) solve() (solved bool) {
	// While frontier is not empty and goal has not been found
	// TODO: This seems dirty, I hate side effects
	for astar.expandFrontNode() != true && astar.frontier.Len() > 0 {
		//bfs.expandFrontier()
	}

	return astar.solved
}

func (astar *AStar) Solution() (path []*Node, err error) {
	if astar.solve() {
		return astar.goal.Path(), nil
	} else {
		return []*Node{}, AStarError{"Solution could not be found for given nodes"}
	}
}

type AStarError struct {
	err string
}

func (p AStarError) Error() string {
	return p.err
}
