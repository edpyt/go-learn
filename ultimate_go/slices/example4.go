package main

import "fmt"

func main() {
  slice := make([]string, 5, 8)

  slice[0] = "Apple"
  slice[1] = "Orange"
  slice[2] = "Banan"
  slice[3] = "Grape"
  slice[4] = "Slave"

  inspectSlice(slice1)

  slice2 := slice1[2:4]
  inspectSlice(slice2)

  fmt.Println("***************")

  slice2[0] = "CHANGED"

  inspectSlice(slice1)
  inspectSlice(slice2)

  fmt.Println("***************")

  slice3 := make([]string, len(slice1))
  copy(slice3, slice1)
  inspectSlice(slice3)
}

func inspectSlice(slc []string)  {
  fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
  for i, s := range slc {
    fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
  }
}
