package alg

import (
	"container/heap"
)

type Ucs struct {
	// Where to start, where to end<
	start *Node
	goal  *Node
  solved bool

	frontier *NodeHeap
}

type NodeHeap []*Node

func (n NodeHeap) Len() int           { return len(n) }
func (n NodeHeap) Less(i, j int) bool { return n[i].PathCost < n[j].PathCost }
func (n NodeHeap) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (n *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*n = append(*n, x.(*Node))
}

func (n *NodeHeap) Pop() interface{} {
	old := *n
	length := len(old)
	x := old[length-1]
	*n = old[0 : length-1]
	return x
}

func NewUcs() (ucs Ucs) {
	ucs.frontier = &NodeHeap{}
	heap.Init(ucs.frontier)
  ucs.solved = false
	return ucs
}

func (ucs *Ucs) SetStart(n *Node) {
	ucs.start = n
  ucs.start.PathCost = 0
	heap.Push(ucs.frontier, n)
}

func (ucs *Ucs) SetGoal(n *Node) {
	ucs.goal = n
}

func (ucs *Ucs) expandFrontNode() (goalReached bool) {
  frontNode := ucs.frontier.Pop().(*Node)

	if ucs.goal == frontNode {
    ucs.solved = true
		return true
	}

	for _, adj := range frontNode.Adj() {
		if adj.PathCost > ( frontNode.PathCost + frontNode.TraversalCost(adj) ) {
			adj.Prev = frontNode
			adj.PathCost = frontNode.PathCost + frontNode.TraversalCost( adj )
      ucs.frontier.Push( adj )
		}
  }
	return false
}

func (ucs *Ucs) solve() (solved bool) {
	// While frontier is not empty and goal has not been found
  // TODO: This seems dirty, I hate side effects
	for ucs.expandFrontNode() != true && ucs.frontier.Len() > 0 {
		//bfs.expandFrontier()
	}

  return ucs.solved
}

func (ucs *Ucs) Solution() (path []*Node, err error) {
	if ucs.solve() {
		return ucs.goal.Path(), nil
	} else {
		return []*Node{}, UcsError{"Solution could not be found for given nodes"}
	}
}


type UcsError struct {
	err string
}

func (p UcsError) Error() string {
	return p.err
}
