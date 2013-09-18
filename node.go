package alg

import (
	"math"
)

// TODO: turn this into an interface

/*
 * Basic building block for all graph traversal algorithms
 */
type Node struct {
	// Contains all adjacent nodes
	adj            []*Node
	adjMap         map[*Node]bool
	X              float64
	Y              float64
	Prev           *Node
	PathCost       float64
	HieuristicCost float64
}

func NewNode() (n Node) {
	n.adj = make([]*Node, 0)
	n.adjMap = make(map[*Node]bool)
	n.PathCost = math.MaxFloat64
	return n
}

func (node *Node) CalcHieuristic(goal *Node) {
	node.HieuristicCost = math.Sqrt(math.Pow(node.X-goal.X, 2) + math.Pow(node.Y-goal.Y, 2))
}

func (node *Node) Adj() (nodes []*Node) {
	return node.adj
}

func (node *Node) AdjTo(adj Node) (adjTo bool) {
	if node.adjMap[&adj] == true {
		return true
	} else {
		return false
	}
}

func (node *Node) AddAdjNode(adjNode *Node) {
	node.adj = append(node.adj, adjNode)
	node.adjMap[adjNode] = true
}

func (n *Node) TraversalCost(node *Node) float64 {
	return math.Sqrt(math.Pow(node.X-n.X, 2) + math.Pow(node.Y-n.Y, 2))
}

func (n *Node) Path() (path []*Node) {
	if n.Prev == nil {
		return []*Node{n}
	} else {
		return append(n.Prev.Path(), n)
	}
}
