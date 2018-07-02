package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func main() {
  /*
    Challenge 1
    This code is a deadlock why?
    Deadlock means that two threads aren't synchronizing, they're waiting
    for each to communicate but they never got notified
  */
  //c := make(chan int)
  //c <- 1 (we have a sender, but no one to receive it)
  //fmt.Println(<-c)

  // Solution to Challenge 1
  // c := make(chan int)
  // go func() {
  //   c <- 1
  //   close(c)
  // }()
  // fmt.Println(<-c)

  /*
    Challenge 2
    Why does this only print 0?
    And what can you do to get it to print all 0-9 numbers?
  */

  // c := make(chan int)
  // go func() {
  //   for i := 0; i < 10; i++ {
  //     c <- i
  //   }
  // }()
  // fmt.Println(<-c)

  // Solution to Challenge 2 Pt.1 (has issues)
  // Q1: Where is the deadlock?
    // In the for loop, it prints value from 0-9
    // However, it's expecting for another goroutine to return
    // But all the goroutines are asleep!
  // Q2: Why are all goroutines asleep?
    // No other channels to read
  // How can we fix this?
    // Using range, waitgroup, or semaphore
  c := make(chan int)
  // go func() {
  //   for i := 0; i < 10; i++ {
  //     c <- i
  //   }
  //
  // }()
  //
  // for {
  //   fmt.Println(<-c)
  // }

  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
    close(c)
  }()

  for n := range c {
    fmt.Println(n)
  }
  /*wg.Add(1)

  go func() {
    for i := 0; i < 10; i++ {
      c <- i
    }
    wg.Done()
  }()

  go func() {
    wg.Wait()
    close(c)
  }()

  for n := range c {
    fmt.Println(n)
  } */
}
