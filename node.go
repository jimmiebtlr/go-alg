package alg

/*
 * Basic building block for all graph traversal algorithms
 */
struct Node interface{
  // Cost to travel to this node
  TraversalCost( node Node )

  // Contains all adjacent nodes
  adj []Node
  adjMap map[]bool
}

func ( node * Node ) Adj()( nodes []Node ){
  return adj
}

func ( node *Node ) AdjTo( node Node )( adjTo bool ){
  if adjMap[ node ] == true {
    return true
  }else{
    return false
  }
}

func ( node *Node ) AddAdjNode( node Node ){
  append( adj, node )
  adjMap[ node ] = true
}

/* 
 * Object to hold path data and value
 */
struct Path type{
  Nodes []Node

  // Is there a better way to do the type on this?
  // Surely
  Cost float64
}

/*
 * Returns the last element in nodes, representing the end of the path
 */
func (p *Path)CurrentNode( )( node Node ){
  return p.Nodes[:len(p.Nodes) - 1 ]
}

func (p *Path)AddNode( node Node )( err error ){
  // Look at last node, if last node has value traveling to node
  // Add it
  p.Cost += p.CurrentNode.TraversalCost( node )
}

func (p *Path)Copy()(path Path ){
  path := Path{
    Nodes: p.Nodes
    Cost: p.Cost
  }
  return path
}
