package main

type data struct {
  name    string
  age     int
}

func (d data) displayName() {
  println("My Name Is", d.name)
}

func (d *data) setAge(age int) {
  d.age = age
  println(d.name, "Is Age", d.age)
}

func main() {
  d := data {
    name: "Bill",
  }

  println("Proper Calls to Methods:")

  d.displayName()
  d.setAge(45)

  println("\nWhat the Compiler is Doing:")

  data.displayName(d)
  (*data).setAge(&d, 45)

  println("\nCall Value Receiver Methods with Variable:")

  f1 := d.displayName

  f1()

  d.name = "Joan"

  f1()

  println("\nCall Pointer Receiver Method with Variable:")

  f2 := d.setAge

  f2(45)
  
  d.name = "Sammy"

  f2(45)
}
