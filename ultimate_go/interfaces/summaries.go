package main

import "fmt"

type reader interface { 
  read(b []byte) (int, error)
}

// Polymorphism
// Can pass any variable in func retrieve, which need interface reader
func retrieve(r reader) error {
  data := make([]byte, 100)

  len, err := r.read(data)
  if err != nil {
    return err
  }

  println(string(data[:len]))
  return nil
}


// Can't use pointer in methods with interfaces
type notifier interface {
  notify()
}

type user struct {
  name    string
  email   string
}

func (u *user) notify() {
  fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
  n.notify()
}

// One more example
type duration int

func (d *duration) notify() {
  fmt.Println("Sending notification in", *d)
}

// Interface slices
type printer interface {
  print()
}

type canon struct {
  name string
}

func (c canon) print() {
  fmt.Printf("Printer Name: %s\n", c.name)
}

type epson struct {
  name string
}

func (e *epson) print() {
  fmt.Printf("Printer Name: %s\n", e.name)
}

func main() {
  // u := user{"Bill", "bill@email.com"}
  // FIXME: 
  // sendNotification(u)

  // FIXME: constants are living only in compilation stage.
  // doesn't have address
  // duration(42).notify()

  c := canon{"PIXMS TR4520"}
  e := epson{"WorkForce Pro WF-3720"}

  printers := []printer{
    c,
    &e,
  }

  c.name = "PROGRAF PRO-100"
  e.name = "Home XP-4100"

  for _, p := range printers {
    p.print()
  }
}
