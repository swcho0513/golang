package main

import "fmt"

func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}

func main() {
  fmt.Println("0! = ", factorial(0))
  fmt.Println("5! = ", factorial(5))
}
