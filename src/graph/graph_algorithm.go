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
func (g *Graph) BellmanFord(src *Node) map[*Node]WeightItem {
  dist := make(map[*Node]WeightItem)
  for _,e := range g.Nodes {
    dist[e] = ^WeightItem(0) // this represents infinity
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

func (g *Graph) BFS(src *Node, fun func(*Node)) {
  q := Queue{}
  seen := make(map[*Node]bool)
  for _,e := range g.Nodes {
    seen[e] = false
  }
  seen[src] = true
  q.Push(src)
  var val *Node
  for !q.IsEmpty() {
    val = q.Pop()
    fun (val)
    /** for all successors **/
    for _,e := range g.Edges[*val] {
      if !seen[e.Vertex] {
        q.Push(e.Vertex)
        seen[e.Vertex] = true
      }
    }
  }
}
