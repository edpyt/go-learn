package main

import (
  "errors"
  "fmt"
  "io"
  "math/rand"
  "time"
)

type Data struct {
  Line string
}

type Puller interface {
  Pull(d *Data) error
}

type Storer interface {
  Store(d *Data) error
}

type PullStorer interface {
  Puller
  Storer
}

type Xenia struct {
  Host string
  Timeout time.Duration
}

func (*Xenia) Pull(d *Data) error {
  switch rand.Intn(10) {
  case 1, 9:
    return io.EOF
  case 5:
    return errors.New("error reading data from Xenia")
  default:
    d.Line = "Data"
    fmt.Println("In:", d.Line)
    return nil
  }
}

type Pillar struct {
  Host    string
  Timeout time.Duration
}

func (*Pillar) Store(d *Data) error {
  fmt.Println("Out:", d.Line)
  return nil
}

type System struct {
  Puller
  Storer
}

func pull(p Puller, data []Data) (int, error) {
  for i := range data {
    if err := p.Pull(&data[i]); err != nil {
      return i, err
    }
  }

  return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
  for i := range data {
    if err := s.Store(&data[i]); err != nil {
      return i, err
    }
  }

  return len(data), nil
}

func Copy(sys *System, batch int) error {
  data := make([]Data, batch)

  for {
    i, err := pull(ps, data)
    if i > 0 {
      if _, err := store(ps, data[:i]); err != nil {
        return err
      }
    }

    if err != nil {
      return err
    }
  }
}

func main() {
  sys := System{
    Puller: &Xenia{
      Host:   "localhost:8000",
      Timeout: time.Second,
    },
    Storer: &Pillar{
      Host:   "localhost:9000",
      Timeout: time.Second,
    },
  }

  if err := Copy(&sys,3); err != io.EOF {
    fmt.Println(err)
  }
}
