/**
** file: graph_test.go
** testing functions for graph
** author: FanatoniQ
**/
package graph

import (
  "testing"
  "fmt"
)

type ExpectedNodeEdges struct {
  nb_nodes int
  nb_edges int
}

func populategraph(g *Graph) ExpectedNodeEdges {
  n1 := Node{"A"}
  n2 := Node{"B"}
  n3 := Node{"C"}
  g.AddNode(&n1)
  g.AddNode(&n2)
  g.AddNode(&n3)
  g.AddEdge(&n1, &n2, 1) // A -1-> B
  g.AddEdge(&n2, &n1, 2) // B -2-> A
  g.AddEdge(&n1, &n3, 4) // A -4-> C
  g.AddEdge(&n2, &n3, 2) // B -2-> C
  return ExpectedNodeEdges{3, 4}
}

func count_edges(g *Graph) int {
  var length int = 0
  for _,e := range(g.Edges) {
    length += len(e)
  }
  return length
}

func TestCreate(t *testing.T) {
  g := InitGraph()
  nb := populategraph(&g)
  if len(g.Nodes) != nb.nb_nodes {
    t.Errorf("Expected %d nodes but only %d present", nb.nb_nodes, len(g.Nodes))
  }
  if length := count_edges(&g); length != nb.nb_edges {
    t.Errorf("Expected %d edges but only %d present", nb.nb_edges, length)
  }
}

func TestDump(t *testing.T) {
  g := InitGraph()
  fmt.Println(g.String())
}
