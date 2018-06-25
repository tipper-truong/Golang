package main

import (
  "fmt"
  "golang-tutorial/src/scope/vis"
  )

var x int = 42

func main() {
  fmt.Println(x)
  foo()
  y := 17
  fmt.Println(y)

  fmt.Println(vis.MyName) // from vis directory
}

func foo() {
  y := 21
  fmt.Println(x)
  fmt.Println(y)
}
