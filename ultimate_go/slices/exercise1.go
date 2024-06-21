package main

func main() {
  x := []int{}

  for i := 1; i < 11; i++ {
    x = append(x, i)
  }

  for _,n := range x {
    println(n)
  }

  println("************")

  y := []string{"bob", "annie", "xiang", "train", "rain"}
  for i, word := range y {
    println(i, word)
  }

  println("************")
   
  s := y[1:2]
  for i, word := range s {
    println(i, word)
  }
}
