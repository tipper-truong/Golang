package main

import (
  "fmt"
)

// Checking for deadlocks
// Checking to make sure main doesn't exit before all processing is done
// unbuffer channels can cause deadlocks?
// Questions to ask yourself:
// Blocking from main executing?
// Am I able to access my channel and receive from my channel?
// Make sure to check for race condition and deadlocks
func main() {
    c1 := incrementor("Foo:")
    c2 := incrementor("Bar:")
    c3 := puller(c1)
    c4 := puller(c2)
    fmt.Println("Final Counter:", <-c3 + <-c4) // inbound channel
}

// Outbound channel
func incrementor(s string) chan int {
  out := make(chan int)
  go func() {
    for i:= 0; i < 20; i++ {
      out <- i
      fmt.Println(s, i)
    }
    close(out)
  }()
  return out
}


func puller(c chan int) chan int {
  out := make(chan int)
  go func() {
    var sum int
    for n := range c {
      sum += n
    }
    out <- sum
    close(out)
  }()
  return out
}
