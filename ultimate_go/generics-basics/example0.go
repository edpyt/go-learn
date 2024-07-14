package main

func genPrint[T any](slice []T) {
  print("Generic: ")

  for _, v := range slice {
    print(v, " ")
  }

  print("\n")
}

func main() {
  numbers := []int{1,2,3}
  genPrint[int](numbers)

  strings := []string{"A", "B", "C"}
  genPrint[string](strings)

  floats := []float64{1.7, 2.2, 3.14}
  genPrint[float64](floats)
}
