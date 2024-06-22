package main

func main() {
  x := make(map[string]int)

  x["bob"] = 1
  x["cola"] = 1
  x["fanta"] = 1
  x["merengi"] = 1
  x["flow"] = 1

  for key, value := range x{
    println(key, value)
  }
}
