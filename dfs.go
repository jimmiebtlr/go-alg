package alg

//import "log"

type Dfs struct {
	// Where to start, where to end<
	start *Node
	goal  *Node
}

func NewDfs() (dfs Dfs) {
	return dfs
}

func (dfs *Dfs) SetStart(n *Node) {
	dfs.start = n
}

func (dfs *Dfs) SetGoal(n *Node) {
	dfs.goal = n
}

func (dfs *Dfs) expandNode(n *Node, visited map[*Node]bool) (nodes []*Node, err error) {
	visited[n] = true
	if n == dfs.goal {
		return []*Node{n}, nil
	}

	for _, node := range n.Adj() {
		if !visited[node] {
			nodes, err = dfs.expandNode(node, visited)
			if err == nil {
				return append(nodes, n), nil
			}
		}
	}
	visited[n] = false
	return []*Node{}, BfsError{"Path not found"}
}

func (dfs *Dfs) solve() (solved bool, path []*Node) {
	path, err := dfs.expandNode(dfs.start, map[*Node]bool{})
	if err == nil {
		return true, path
	} else {
		return false, []*Node{}
	}
}

func reverse(path []*Node) (rev_path []*Node) {
	rev_path = make([]*Node, len(path))
	for i := 0; i < len(path)/2; i++ {
		tmp_n := path[i]
		path[i] = path[len(path)-i-1]
		path[len(path)-i-1] = tmp_n
	}
	return rev_path
}

func (dfs *Dfs) Solution() (path []*Node, err error) {
	solved, path := dfs.solve()
	reverse(path)
	if solved {
		return path, nil
	} else {
		return []*Node{}, DfsError{"Solution could not be found for given nodes"}
	}
}

type DfsError struct {
	err string
}

func (p DfsError) Error() string {
	return p.err
}
