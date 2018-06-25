package main

import "fmt"

const (
  A = iota
  B = iota
  C = iota << 2
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
