package main

type Toy struct {
  Name   string
  Weight float64

  onHand string
  sold   string
}

func New(name string, weight float64, onHand string, sold string) Toy {
  return Toy{name, weight, onHand, sold}
}

func (t Toy) OnHand() string {
  return t.onHand
}

func (t Toy) Sold() string {
  return t.sold
}

// import toy

func main() {
  t := New("bob", 1.23, "yes", "no")

  println(t.OnHand())
  println(t.Sold())
}
