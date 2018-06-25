package main

import "fmt"
import "strconv"

func main() {
  fmt.Println("Enter your name: ")
  input := ""
  fmt.Scan(&input)
  fmt.Println("Hello " + input)

  fmt.Println("Enter a small number: " )
  num1 := ""
  fmt.Scan(&num1)

  fmt.Println("Enter a large number: " )
  num2 := ""
  fmt.Scan(&num2)

  x, _ := strconv.Atoi(num2)
  y, _ := strconv.Atoi(num1)

  if x < y {
    fmt.Println("Incorrect Input")
  } else if (x == y) {
    fmt.Println("They're not suppose to be equal")
  } else {
    var res int = x % y
    fmt.Printf("%d\n", res)
  }

}
