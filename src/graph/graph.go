package graph

import(
  "fmt"
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

// printing methods

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

// functionnal methods

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

// other methods

func Min(a, b uint) uint {
  if a > b { return b } else { return a }
}

// algorithms

func (g *Graph) BellmanFord(src *Node) map[*Node]uint {
  dist := make(map[*Node]uint)
  for _,e := range g.Nodes {
    dist[e] = ^uint(0) // this represents infinity
  }
  dist[src] = 0
  for k := 0; k < len(g.Nodes); k++ {
    for _,e := range(g.Edges[*g.Nodes[k]]) {
      // for each graph edge (*g.Nodes[k], e)
      dist[e.Vertex] = Min(dist[e.Vertex], dist[g.Nodes[k]] + e.Weight)
    }
  }
  return dist
}
