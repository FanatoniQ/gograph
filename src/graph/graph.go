/**
** file: graph.go
** graph golang datastructure
** author: FanatoniQ
**/
package graph

import(
)

/** generic item type **/
type Item interface{}

/** Node data structure **/
type Node struct {
  Value Item
}

/** Edge representation
** store weight to go to Vertex which is the destination
**/
type Edge struct {
  Vertex *Node
  Weight uint
}

/** Graph includes a list of allocated nodes
** and a map to map a node to an edge
**/
type Graph struct {
  Nodes []*Node /** store our Graph nodes **/
  Edges map[Node][]*Edge
  /**
  ** each node has a list of edges
  **/
}

func (g *Graph) AddNode(n *Node) {
  g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(src, dst *Node, w uint) {
  g.Edges[*src] = append(g.Edges[*src], &Edge{dst, w})
}

func InitGraph() Graph {
  g := Graph{}
  g.Edges = make(map[Node][]*Edge)
  return g
}
