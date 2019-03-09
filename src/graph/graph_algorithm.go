/**
** file: graph_algorithm.go
** graph golang graph algorithms
** such as :
**    - BellmanFord
** author: FanatoniQ
**/
package graph

import(
)

/**
** BellmanFord algorithm
** return the map of Nodes and distance to the given source node src
**/
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
