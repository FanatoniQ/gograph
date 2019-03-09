/**
** file: queue_test.go
** testing functions for queues
** author: FanatoniQ
**/
package graph

import (
  "testing"
)

func TestFunctional(t *testing.T) {
  q := Queue{}
  var lst []*Node
  lst = append(lst, &Node{"A"})
  lst = append(lst, &Node{"B"})
  lst = append(lst, &Node{"C"})
  lst = append(lst, &Node{"D"})
  lst = append(lst, &Node{"E"})
  for i,_ := range lst {
    q.Push(lst[i])
  }
  var i int = len(lst) - 1
  for !q.IsEmpty() {
    if q.Pop() != lst[i] {
      t.Errorf("Wrong queue behaviour !")
    }
    i--
  }
  if i != -1 {
    t.Errorf("%d elements not added to queue !", i + 1)
  }
}
