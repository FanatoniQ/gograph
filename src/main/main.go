package main

import (
    "fmt"
    "graph"
)

func toString(n *graph.Node) {
  fmt.Println(n)
}

func main() {
  g := graph.InitGraph()
  n1 := graph.Node{"A"}
  n2 := graph.Node{"B"}
  n3 := graph.Node{"C"}
  g.AddNode(&n1)
  g.AddNode(&n2)
  g.AddNode(&n3)
  g.AddEdge(&n1, &n2, 1) // A -1-> B
  g.AddEdge(&n2, &n1, 2) // B -2-> A
  g.AddEdge(&n1, &n3, 4) // A -4-> C
  g.AddEdge(&n2, &n3, 2) // B -2-> C

  dist := g.BellmanFord(&n1)
  for i,e := range(dist) {
    fmt.Println(i, e)
  }

  fmt.Println(g.String())

  g.BFS(&n1, toString)
}
