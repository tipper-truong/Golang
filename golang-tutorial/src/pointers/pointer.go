package main

import "fmt"

func main() {
  a := 43

  fmt.Println(a) // 43
  fmt.Println(&a) // get the memory address

  var b *int = &a
  fmt.Println(b) // memory address
  fmt.Println(*b) // 43

  *b = 10
  fmt.Println(a) // prints 10 b/c you have the memory address of a in b.
}
