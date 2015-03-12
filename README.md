cqueue
======
> circular buffer implementation in golang using linked-list

[![GoDoc](https://godoc.org/github.com/pravj/cqueue?status.svg)](http://godoc.org/github.com/pravj/cqueue)

#####Installation
```
go get github.com/pravj/cqueue
```

> A circular buffer structure
```go
// CQueue represents a circular buffer data structure
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
  // middleware function to call when buffer overflows
  overflow func()
}
```

#####[Package API](http://godoc.org/github.com/pravj/cqueue)
* Accessor Methods
  * Value
  * IsFull
  * IsEmpty
* Mutator Methods
  * Push
  * Pop

#####Usage
```go
import (
  "fmt"
  "github.com/pravj/cqueue"
)

// Alert is the middleware function to call
// in case of buffer overflow
//
// Leave its content blank if you don't want a middleware
func Alert() {
  fmt.Println("Buffer Overflow")
}

func main() {
  args := map[string]interface{}{"size":3, "middleware": Alert}
  cq := cqueue.New(args) // [- - -]
  
  cq.Push(4) // [4 - -]
  cq.Push(4) // [4 4 -]
  cq.Push(5) // [4 4 5]
  cq.Push(6) // [4 5 6] will invoke the middleware function also
  cq.Pop()   // [5 6 -]
  cq.Push(7) // [5 6 7]
  cq.Pop()   // [6 7 -]
  cq.Pop()   // [7 - -]
  cq.Pop()   // [- - -]
  cq.Pop()   // error "cqueue: buffer is empty"
}
```

#####TODO

- [ ] Test Suite
- [ ] Distributed Locking on buffer operations

---

Made by [Pravendra Singh](http://pravj.github.io)
