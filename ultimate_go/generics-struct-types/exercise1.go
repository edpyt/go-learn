package main

import "errors"

type stack[T any] struct {
  data []T
}

func (s *stack[T]) push(v T) {
  s.data = append(s.data, v)
}

func (s *stack[T]) pop() (T, error) {
  var zero T

  if len(s.data) == 0 {
    return zero, errors.New("stack is empty")
  }
  
  el := s.data[len(s.data) -1]
  s.data = s.data[:len(s.data) - 1]
  return el, nil  
}

func main() {
  var s stack[int]

  s.push(10)

  println(s.data)

  v, _ := s.pop()

  println(v)

  v, _ = s.pop()

  println(v)

  _, err := s.pop()
  println(err)
}
