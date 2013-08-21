package main

import(
  "jimmiebtlr/alg"
  "math"
)

struct BasicNode type{
  x float64
  y float64
}

func (n *BasicNode)TraversalCost( node *alg.Node ){
  return math.Pow(node.X -n.X,2) + math.Pow( node.Y - n.Y ,2 )
}

func main(){
  nodes := {
    Node

  //bfs := Bfs{}
}
