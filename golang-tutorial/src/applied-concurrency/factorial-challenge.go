package main

import (
  "fmt"
)

// Challenge 1: Use goroutines and channels to calculate factorial
/*
  Challenge 2:
  Why might you want to use goroutines and channels to calculate factorial?

  You want to use goroutines and channels to calculate factorial because
  A goroutine, concurrency, and paralleism is helpful when you have a lot of
  processing to do. If I needed to process thousands of factorial calculations,
  then putting them into goroutine and running those calculations concurrently
  and in parallel - using all of the CPU cores on my machine, will help those calculations
  get done faster

*/
func main() {
  //f := factorial(4)
  factorialChan := factorial(4)
  for n := range factorialChan { // range blocks from exiting main
    fmt.Println("Result:", n)
  }
}

func factorial(n int) chan int {
  out := make(chan int)

  go func() {
    total := 1
    for i := 1; i <= n; i++ {
      total *= i
    }
    out <- total
    close(out)
  }()

  return out
}
