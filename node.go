package alg

import (
  "math"
)

/*
 * Basic building block for all graph traversal algorithms
 */
type Node struct{
  // Contains all adjacent nodes
  adj []*Node
  adjMap map[*Node]bool
  X float64
  Y float64
}

func NewNode()( n Node ){
  n.adj = make([]*Node, 0)
  n.adjMap = make( map[*Node]bool )
  return n
}

func ( node *Node ) Adj()( nodes []*Node ){
  return node.adj
}

func ( node *Node ) AdjTo( adj Node )( adjTo bool ){
  if node.adjMap[ &adj ] == true {
    return true
  }else{
    return false
  }
}


func ( node *Node ) AddAdjNode(  adjNode *Node ){
  node.adj = append( node.adj, adjNode )
  node.adjMap[ adjNode ] = true
}

func (n *Node)TraversalCost( node *Node )( float64 ){
  return math.Sqrt(math.Pow(node.X -n.X,2) + math.Pow( node.Y - n.Y ,2 ))
}

/* 
 * Object to hold path data and value
 */
type Path struct{
  Nodes []*Node

  // Is there a better way to do the type on this?
  // Surely
  Cost float64
}

/*
 * Returns the last element in nodes, representing the end of the path
 */
func (p *Path)CurrentNode( )( node *Node, err error ){
  if len( p.Nodes ) > 0 {
    return p.Nodes[ len(p.Nodes) - 1 ], nil
  }else{
    return &Node{}, &pathError{"Path has no nodes" }
  }
}

type pathError struct {
    s string
}

func (e *pathError) Error() string {
    return e.s
}

func (p *Path)AddNode( node *Node ){
  // Look at last node, if last node has value traveling to node
  // Add it
  cn,err := p.CurrentNode()
  if err == nil {
    p.Cost += cn.TraversalCost( node )
  }
  p.Nodes = append( p.Nodes, node )
}

func (p *Path)Copy()(path Path ){
  nodes := make([]*Node, len( p.Nodes ) )
  copy( nodes, p.Nodes )
  path = Path{
    Nodes: nodes,
    Cost: p.Cost,
  }
  return path
}
