package main

import (
  "testing"
	"jimmiebtlr/alg"
)

func TestBFS(t *testing.T) {
	n1 := alg.NewNode()
	n1.X = 1
	n1.Y = 2

	n2 := alg.NewNode()
	n2.X = 3
	n2.Y = 3

	n3 := alg.NewNode()
	n3.X = 4
	n3.Y = 4

	n4 := alg.NewNode()
	n4.X = 5
	n4.Y = 5

	n5 := alg.NewNode()
	n5.X = 10
	n5.Y = 20

	n1.AddAdjNode(&n2)
	n2.AddAdjNode(&n3)
	n3.AddAdjNode(&n4)
	n4.AddAdjNode(&n5)

	bfs := alg.NewBfs()
	bfs.SetStart( &n1 )
	bfs.SetGoal( &n5 )
  solution := bfs.Solution()

  if solution.Nodes[0] != &n1 {
    t.Errorf( "The first step in the path was not correct" )
  }
  if solution.Nodes[1] != &n2 {
    t.Errorf( "The second step in the path was not correct" )
  }
  if solution.Nodes[2] != &n3 {
    t.Errorf( "The third step in the path was not correct" )
  }
  if solution.Nodes[3] != &n4 {
    t.Errorf( "The fourth step in the path was not correct" )
  }
  if solution.Nodes[4] != &n5 {
    t.Errorf( "The fifth step in the path was not correct" )
  }
}
