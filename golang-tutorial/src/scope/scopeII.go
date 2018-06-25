package main
import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
  x:= 10
  y:= 20
  sum := add(x, y)
  fmt.Println(sum)
}
