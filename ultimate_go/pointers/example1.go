package main

func main() {
  count := 10

  println("count:\tValue Of [", count, "]\t Addr Of [", &count, "]")
  increment1(count)
  println("count:\tValue Of [", count, "]\tAddr Of[", &count, "]")

  println("count:\tValue Of [", count, "]\t Addr Of [", &count, "]")
  increment2(&count)
  println("count:\tValue Of [", count, "]\tAddr Of[", &count, "]")
}

func increment1(inc int) {
  inc++
  println("inc1:\tValue Of [", inc, "]\tAddr Of[", &inc, "]")
}

func increment2(inc *int) {
  *inc++
  println("inc2:\tValue Of [", inc, "]\tAddr Of[", &inc,
  "]\tPoints To[", *inc, "]")
}
