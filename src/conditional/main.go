package main 

import "fmt"

func main() {
  x := 10
  
  // No parenthesis, as opposed to JS which does use parenthesis
  // Go also has `else if` and `else`  
  if x <= 10 {
    fmt.Println("Equal to or under 10")
  }
}
