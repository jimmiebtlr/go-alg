package main

import (
	"jimmiebtlr/alg"
	"log"
)

func main() {
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

  log.Println( solution )
}
