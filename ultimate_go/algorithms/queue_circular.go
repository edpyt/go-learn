package main

import (
  "errors"
  "fmt"
)

func main() {
  const items = 5

  q, err := New(items)
  if err != nil {
    fmt.Println(err)
    return 
  }

  for i := 0; i < items; i++ {
    name := fmt.Sprintf("Name%d", i)
    if err := q.Enqueue(Data{Name: name}); err != nil {
      fmt.Println(err)
      return
    }

    fmt.Println("queue:", name)
  }

  fmt.Println("-----------------")

  f := func(d Data) error {
    fmt.Println("enqueue:", d.Name)
    return nil
  }

  q.Operate(f)
}

type Data struct {
  Name string
}

type Queue struct {
  Count   int
  data    []Data
  front   int
  end     int
}

func New(cap int) (*Queue, error)  {
  if cap <= 0 {
    return nil, errors.New("invalid capacity")
  }

  q := Queue{
    front: 0,
    end:   0,
    data:  make([]Data, cap),
  }
  return &q, nil
}

func (q *Queue) Enqueue(data Data) error {
  if q.front+1 == q.end || q.front == len(q.data) && q.end == 0  {
    return errors.New("queue at capacity")
  }

  switch {
  case q.front == len(q.data):
    q.front = 0
    q.data[q.front] = data
  default:
    q.data[q.front] = data
    q.front++
  }

  q.Count++

  return nil
}

func (q *Queue) Dequeue() (Data, error) {
  if q.front == q.end {
    return Data{}, errors.New("qeuue is empty")
  }

  var data Data
  switch {
  case q.end == len(q.data):
    q.end = 0
    data = q.data[q.end]
  default:
    data = q.data[q.end]
    q.end++
  }

  q.Count--

  return data, nil
}

func (q *Queue) Operate(f func(d Data) error) error {
  end := q.end
  for {
    if end == q.front {
      break
    }

    if end == len(q.data) {
      end = 0
    }

    if err := f(q.data[end]); err != nil {
      return err
    }

    end++
  }
  return nil
}
