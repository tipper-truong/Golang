package main
import "fmt"

func wrapper() func() int { // returns a function
  x := 0

  return func() int {
    x++
    return x
  }
}

func main() {
  increment := wrapper() // gets value from inner scope of return function
  fmt.Println(increment())
  fmt.Println(increment())
}
