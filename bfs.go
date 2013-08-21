package alg

struct Bfs type{
  // Where to start, where to end<
  Start Node
  Goal Node

  visited map[]bool
  frontier []Path
}

func ( bfs *Bfs )expandFrontier( ) {
  for _,path := range frontier {
    for _,adj := range path.CurrentNode.Adj {
      if visited[adj] != true {
        p := path.Copy()
        p.AddNode( adj )
        visited[ adj ] = true
      }
    }
    frontier.Remove( path )
  }
}

func (bfs *Bfs)Solve(){
  // While frontier is not empty and goal has not been found
  for len(bfs.frontier) != 0 && !bfs.goalReached {
    // Expand the current frontier
    bfs.expandFrontier
  }
}

func (bfs *Bfs)goalReached()( goal bool ){
  for _,path := range frontier {
    if path.CurrentNode == bfs.Goal {
      return true
    }
  }
  return false
}
