/**
** file: graph_printing.go
** printing functions for graph
** author: FanatoniQ
**/
package graph

import(
  "fmt"
)

func (n *Node) String() string {
  return fmt.Sprintf("%v", n.Value)
}

func (e *Edge) String() string {
  return e.Vertex.String() + " : " + fmt.Sprint(e.Weight)
}

func (g *Graph) String() string {
  var dump string = ""
  for _,s := range g.Nodes {
    dump += s.String() + " ----> "
    for _,e := range g.Edges[*s] {
      dump += "(" + e.String() + ") "
    }
    dump += "\n"
  }
  return dump
}
