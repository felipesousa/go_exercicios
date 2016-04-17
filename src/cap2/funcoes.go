package main

import("fmt")

func soma(a,b int) int {
  c := a + b
  return c
}

func main() {
  d := soma(13, 42)
  fmt.Printf("%d", d)
}
