package main

import "fmt"

func main() {
  // `:=` is a shorthand for `var <name> <type> = <value>`
  x := 5
  y := 10
  sum := x + y
  
  fmt.Println("Sum of x plus y:")
  fmt.Println(sum)
}
