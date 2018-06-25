package main
import "fmt"

func run_loop(i *int) {
  for *i < 10 {
    fmt.Println(*i)
    *i++
  }
}
func main() {
  for i := 0; i < 100; i++ {
    fmt.Println(i)
  }

  i := 0
  run_loop(&i)

}
