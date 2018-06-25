package main

import "fmt"

func main() {
  x := 42
  fmt.Println(x)
  { // inner scope
    x := 52
    fmt.Println(x)
    y := "Credit belongs with the one who is in the ring."
    fmt.Println(y)
  }
}
