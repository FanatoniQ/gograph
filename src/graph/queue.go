/**
** file: queue.go
** queue implementation in golang
** author: FanatoniQ
**/
package graph

import(
)

/** Queue datastructure **/
type Queue struct {
  vals []*Node /** we just have a simple vector for our values **/
}

/**
** returns true when the queue is empty
**/
func (q *Queue) IsEmpty() bool {
  return len(q.vals) == 0
}

/**
** Push a QItem to the queue
**/
func (q *Queue) Push(val *Node) {
  q.vals = append(q.vals, val)
}

/**
** Pop the last Pushed item from the Queue
** panic() when called on empty queue
**/
func (q *Queue) Pop() *Node {
  if q.IsEmpty() {
    panic("Should not be empty")
  }
  length := len(q.vals) - 1
  ret := q.vals[length]
  q.vals = q.vals[:length]
  return ret
}
