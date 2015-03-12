// Package cqueue implement circular buffer/queue data structure
package cqueue

import (
  "container/list"
  "fmt"
)

// TODO: overflow middleware function also

type CQueue struct {
  size int
  Length int

  Queue *list.List
  Head *list.Element
  Tail *list.Element
}

func New(size int) *CQueue {
  return &CQueue{size: size, Queue: list.New()}
}

func (cq *CQueue) Push(element interface{}) {
  if (cq.Length < cq.size) {
    cq.Tail = cq.Queue.PushBack(element)
    if (cq.Length == 0) {
      cq.Head = cq.Tail
    }

    cq.updateLength(1)
  } else {
    fmt.Println("full buffer")
    cq.Pop()
    cq.Push(element)
  }
}

func (cq *CQueue) Pop() {
  if (cq.Length > 0) {
    head := cq.Head.Next()
    cq.Queue.Remove(cq.Head)
    cq.Head = head

    cq.updateLength(-1)
  } else {
    fmt.Println("empty buffer")
  }
}

func (cq *CQueue) isFull() bool {
  return (cq.Length == cq.size)
}

func (cq *CQueue) isEmpty() bool {
  return (cq.Length == 0)
}

func (cq *CQueue) updateLength(update int) {
  cq.Length += update
}
