package main

import (
  "fmt"
)


/*
  Challenge 1:
  - Change the code to execute 100 factorial computations concurrently
  and in parallel
  - Use the "pipeline" pattern to accomplish this

  What realizations did you have about working with concurrent code when building your solution?
  eg: what insights, did you have which helped you understand working with concurrency

*/
func main() {

  in := gen()

  f := factorial(in)

  for n := range f {
    fmt.Println(n)
  }
}

func gen() <-chan int {
  out := make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
      for j := 3; j < 13; j++ {
        out <- j
      }
    }
    close(out)
  }()
  return out
}

func factorial(in <-chan int) <-chan int {
  out := make(chan int)

  go func() {
    for n := range in {
      out <- fact(n)
    }
    close(out)
  }()

  return out
}

func fact(n int) int {
  total := 1
  for i := 1; i <= n; i++ {
    total *= i
  }
  return total
}
