package main

func event(message string) {
  println(message)
}

type data struct {
  name    string
  age     int
}

func (d *data) event(message string) {
  println(d.name, message)
}

func fireEvent1(f func(string)) {
  f("anonymous")
}

type handler func(string)

func fireEvent2(h handler) {
  h("handler")
}

func main() {
  d := data{
    name: "Bill",
  }

  fireEvent1(event)
  fireEvent1(d.event)

  fireEvent2(event)
  fireEvent2(d.event)

  h1 := handler(event)
  h2 := handler(d.event)

  fireEvent2(h1)
  fireEvent2(h2)

  fireEvent1(h1)
  fireEvent1(h2)
}
