package alg

import (
  "log"
)

type Bfs struct {
	// Where to start, where to end<
	start *Node
	goal  *Node

	visited  map[*Node]bool
	frontier []Path
  solution Path
}

func NewBfs() (bfs Bfs) {
	bfs.visited = make(map[*Node]bool)
	bfs.frontier = make([]Path, 0)
	return bfs
}

func( bfs *Bfs ) SetStart( n *Node ){
  bfs.start = n
  bfs.frontier = []Path{ Path{ []*Node{n},0 } }
}

func( bfs *Bfs ) SetGoal( n *Node ){
  bfs.goal = n
}

func (bfs *Bfs) expandFrontier() {
	newFrontier := make([]Path, 0)

  // each path in the frontier
	for _, path := range bfs.frontier {
    // Select the last node from the path
		frontNode, err := path.CurrentNode()
		if err == nil {
      // add each node that is adjacent to the frontier
			for _, adj := range frontNode.Adj() {
				//if bfs.visited[&adj] != true {
					p := path.Copy()
					p.AddNode(adj)
					bfs.visited[adj] = true
					newFrontier = append(newFrontier, p)
				//}
			}
    }
	}
	bfs.frontier = newFrontier
}

func (bfs *Bfs) solve() {
	// While frontier is not empty and goal has not been found
	for len(bfs.frontier) > 0 && !bfs.goalReached() {
		// Expand the current frontier
		bfs.expandFrontier()
	}
}

// TODO get rid of side effect
func (bfs *Bfs) goalReached() (goal bool) {
	for _, path := range bfs.frontier {
		node, err := path.CurrentNode()
		if err == nil {
			if node == bfs.goal {
        log.Println( "Solution found" )
        bfs.solution = path
				return true
			}
		}
	}
	return false
}

func (bfs *Bfs) Solution()(p Path){
  bfs.solve()
  return bfs.solution
}
