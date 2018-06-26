package main

import "fmt"

func main() {
  data := []float64{43, 56, 87, 12, 45, 57} // slice
  n := average(data...) // variadic arguments
  fmt.println(n)
}

func average(sf ...float64) float64 {
  fmt.Println(sf)
  fmt.Printf("%T\n", sf)
  total := 0.0
  for _, v := range sf { // for _, v --> for index, value. We don't want to print index so that's why _
    total += v
  }
  return total / float64(len(sf))
}
