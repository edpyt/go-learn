package main

import (
  "fmt"
  "strings"
)

func main() {
  var l List

  const nodes = 5
  for i := 0; i < nodes; i++ {
    data := fmt.Sprintf("Node%d", i)
    l.Add(data)
  }

  f := func(n *Node) error {
    fmt.Println("Data:", n.Data)
    return nil
  }

  l.Operate(f)

  fmt.Println("---------------------")

  l.OperateReverse(f)
}

type Node struct {
  Data string
  next *Node
  prev *Node
}

type List struct {
  Count int
  first *Node
  last  *Node
}

func (l *List) Add(data string) *Node {
  n := Node{
    Data: data,
    prev: l.last,
  }

  l.Count++

  if l.first == nil && l.last == nil {
    l.first = &n
    l.last = &n
    return &n
  }

  l.last.next = &n
  l.last = &n

  return &n
}

func (l *List) AddFont(data string) *Node {
  n := Node{
    Data: data,
    next: l.first,
  }

  l.Count++

  if l.first == nil && l.last == nil {
    l.first = &n
    l.last = &n
    return &n
  }

  l.first.prev = &n

  l.first = &n

  return &n
}

func (l *List) Find(data string) (*Node, error) {
  n := l.first
  for n != nil {
    if n.Data == data { 
      return n, nil
    }
    n = n.next
  }
  return nil, fmt.Errorf("unable to locate %q in list", data)
}

func (l *List) Remove(data string) (*Node, error) {
  n, err := l.Find(data)
  if err != nil {
    return nil, err
  }

  n.prev.next = n.next
  n.next.prev = n.prev
  l.Count--

  return n, nil
}

func (l *List) Operate(f func(n *Node) error) error {
  n := l.first

  for n != nil {
    if err := f(n); err != nil {
      return err
    }
    n = n.next
  }
  return nil
}

func (l *List) OperateReverse(f func(n *Node) error) error {
  n := l.last

  for n != nil {
    if err := f(n); err != nil {
      return err
    }
    n = n.prev
  }
  return nil
}

func (l *List) AddSort(data string) *Node {
  if l.first == nil {
    return l.Add(data)
  }

  n := l.first
  for n != nil {
    if strings.Compare(data, n.Data) > 0 {
      n = n.next
      continue
    }

    new := Node{
      Data: data,
      next: n,
      prev: n.prev,
    }

    l.Count++

    if l.first == n {
      l.first = &new
    }

    if n.prev != nil {
      n.prev.next = &new
    }

    n.prev = &new

    return n
  }

  return l.Add(data)
}
