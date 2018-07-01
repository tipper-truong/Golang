package main

import (
  "fmt"
)

func main() {

  c := make(chan int)

  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
    close(c) // no longer put values on a channel
  }()

  for n := range c {
    fmt.Println(n)
  }
}
