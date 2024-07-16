package main

import "fmt"

type node[T any] struct {
  Data T
  next *node[T]
  prev *node[T]
}

type list[T any] struct {
  first *node[T]
  last  *node[T]
}

func (l *list[T]) add(data T) *node[T] {
  n := node[T]{
    Data: data,
    prev: l.last,
  }

  if l.first == nil {
    l.first = &n
    l.last  = &n
    return &n
  }

  l.last.next = &n
  l.last      = &n
  return &n
}

type user struct {
  name string
}

func main() {
  var lv list[user]

  n1 := lv.add(user{"bill"})
  n2 := lv.add(user{"ale"})
  fmt.Println(n1.Data, n2.Data)

  var lp list[*user]
  n3 := lp.add(&user{"bill"})
  n4 := lp.add(&user{"ale"})
  fmt.Println(n3.Data, n4.Data)
}
