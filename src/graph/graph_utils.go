/**
** file: graph_utils.go
** utility functions for graph
** author: FanatoniQ
**/
package graph

import(
)

func Min(a, b WeightItem) WeightItem {
  if a > b { return b } else { return a }
}
