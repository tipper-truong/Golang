package main

import "fmt"

func printEven() {
  for i := 0; i <= 100; i++ {
    if(i % 2 == 0) {
      fmt.Println(i)
    }
  }
}

func fizzBuzz() {
  for i:=3; i < 100; i++ {
    if (i % (3 * 5) == 0) {
     fmt.Println("FizzBuzz")
    } else if (i % 3 == 0) {
      fmt.Println("Fizz")
    } else if (i % 5 == 0) {
      fmt.Println("Buzz")
    }
  }
}

func multiplesOf3and5() {
  sum := 0
  for i := 3; i < 1000; i++ {
    if (i % 3 == 0 || i % 5 == 0) {
      sum += i
    }
  }
  fmt.Println(sum)
}

func half(x int) (int, bool) {
  if(x % 2 == 0) {
    return x, true
  }
  return x, false
}

func main() {
  //printEven()

  //fizzBuzz()

  //multiplesOf3and5()

  a, b := half(2)
  fmt.Printf("(%d, ", a)
  fmt.Printf("%t)\n",  b)
}
