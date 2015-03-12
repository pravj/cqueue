// Package cqueue implement circular buffer/queue operations
package cqueue

import (
  "container/list"
  "errors"
)

// CQueue represents a circular buffer data structure
// TODO: overflow middleware function also
type CQueue struct {
  // size limit of the buffer
  size int
  // occupied size of the buffer
  Length int

  // list containing all the buffer elements
  Queue *list.List
  // element at the front-side of the buffer
  Head *list.Element
  // element at the back-side of the buffer
  Tail *list.Element
}

// New function returns a CQueue instance of custom size
func New(size int) *CQueue {
  return &CQueue{size: size, Queue: list.New()}
}

// Push adds an element to the circular buffer
func (cq *CQueue) Push(element interface{}) (el *list.Element, err error) {
  if (cq.Length < cq.size) {
    cq.Tail = cq.Queue.PushBack(element)
    if (cq.Length == 0) {
      cq.Head = cq.Tail
    }

    cq.updateLength(1)
  } else {
    cq.Pop()
    cq.Push(element)
  }

  return cq.Tail, nil
}

// Pop removes the element at the front in circular buffer
// It returns the popped element and error, if any.
func (cq *CQueue) Pop() (el *list.Element, err error) {
  if (cq.Length > 0) {
    el, err = cq.Head, nil
    newHead := cq.Head.Next()

    cq.Queue.Remove(cq.Head)
    cq.Head = newHead
    if (cq.Length == 1) {
      cq.Tail = nil
    }

    cq.updateLength(-1)
  } else {
    el, err = nil, errors.New("cqueue: buffer is empty")
  }

  return el, err
}

// Value returns the value of a list element
func (cq *CQueue) Value(element *list.Element) interface{} {
  if (element == nil) {
    return nil
  }
  return element.Value
}

// Whether the circular buffer is full or not
func (cq *CQueue) isFull() bool {
  return (cq.Length == cq.size)
}

// Whether the circular buffer is empty or not
func (cq *CQueue) isEmpty() bool {
  return (cq.Length == 0)
}

// Updates the occupied length of the circular buffer
func (cq *CQueue) updateLength(update int) {
  cq.Length += update
}
