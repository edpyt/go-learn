package main

func main() {
  var fruits [5]string
  fruits[0] = "Apple"
  fruits[1] = "Orange"
  fruits[2] = "Banana"
  fruits[3] = "Grape"
  fruits[4] = "Plum"

  for i, fruit := range fruits {
    println(i, fruit)
  }

  numbers := [4]int{10,20,30,40}

  for i := 0; i < len(numbers); i++ {
    println(i, numbers[i])
  }
}
