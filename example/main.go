package main

import "fmt"
import "github.com/pravj/cqueue"

func Alert() {
  fmt.Println("Buffer Overflow")
}

func main() {
  args := map[string]interface{}{"size":3, "middleware": Alert}

  cq := cqueue.New(args)
  fmt.Println(cq.Value(cq.Head), "back")
  fmt.Println(cq.Value(cq.Tail), "front")
  fmt.Println(cq.Length)

  cq.Push(4)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Push(4)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Push(5)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Push(6)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Pop()
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Push(7)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Pop()
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Pop()
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Pop()
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  _, err := cq.Pop()
  fmt.Println(err)

  fmt.Println(cq.Value(cq.Head), "back")
  fmt.Println(cq.Value(cq.Tail), "front")
  fmt.Println(cq.Length)

  cq.Push(4)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Push(4)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)

  cq.Push(5)
  fmt.Println(cq.Value(cq.Head), "front")
  fmt.Println(cq.Value(cq.Tail), "back")
  fmt.Println(cq.Length)
}
