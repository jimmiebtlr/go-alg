package alg

//import "log"

type Bfs struct {
	// Where to start, where to end<
	start *Node
	goal  *Node

	frontier []*Node
}

func NewBfs() (bfs Bfs) {
	bfs.frontier = make([]*Node, 0)
	return bfs
}

func (bfs *Bfs) SetStart(n *Node) {
	bfs.start = n
	bfs.frontier = []*Node{n}
}

func (bfs *Bfs) SetGoal(n *Node) {
	bfs.goal = n
}

func (bfs *Bfs) expandFrontier() (goalReached bool) {
	newFrontier := make([]*Node, 0)

	// each path in the frontier
	for _, frontNode := range bfs.frontier {
		// add each node that is adjacent to the frontier
		for _, adj := range frontNode.Adj() {
			if adj.Prev == nil {
				adj.Prev = frontNode
				newFrontier = append(newFrontier, adj)
			}

			if adj == bfs.goal {
				return true
			}
		}
	}
	bfs.frontier = newFrontier
	return false
}

func (bfs *Bfs) solve() (solved bool) {
	// While frontier is not empty and goal has not been found
	for bfs.expandFrontier() != true && len(bfs.frontier) > 0 {
		//bfs.expandFrontier()
	}

	// kind of hackish but works for now
	if len(bfs.frontier) == 0 {
		return false
	} else {
		return true
	}
}

func (bfs *Bfs) Solution() (path []*Node, err error) {
	if bfs.solve() {
		return bfs.goal.Path(), nil
	} else {
		return []*Node{}, BfsError{"Solution could not be found for given nodes"}
	}
}

type BfsError struct {
	err string
}

func (p BfsError) Error() string {
	return p.err
}
