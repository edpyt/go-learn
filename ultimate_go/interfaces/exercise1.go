package main

type speaker interface {
  speak() 
}

type english struct {}
type chinese struct {}

func (e english) speak() {
  println("Hello World")
}

func (c chinese) speak() {
  println("你好世界")
}

func main() {
  var s speaker;

  s = english{};
  s.speak();

  s = chinese{};
  s.speak();

  e := english{};
  c := chinese{};

  sayHello(e);
  sayHello(c);
}

func sayHello(sp speaker) {
  sp.speak()
}
