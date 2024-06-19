package main

import "fmt"

func main() {
  fruits := make([]string, 5, 8)
  fruits[0] = "Apple"
  fruits[1] = "Orange"
  fruits[2] = "Banan"
  fruits[3] = "Grape"
  fruits[4] = "Slave"

  inspectSlice(fruits)

}

func inspectSlice(slice []string) {
  fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
  for i, s := range slice {
    fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
  }
}
