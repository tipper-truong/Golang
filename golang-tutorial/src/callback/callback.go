package main

import "fmt"

// pass 1,2,3,4, callback and prints each number
func visit(numbers []int, callback func(int)) {
    for _, n := range numbers {
      callback(n)
    }
}

func main() {
  visit([]int{1,2,3,4}, func (n int) {
    fmt.Println(n)
  })
}
