package main

import (
  "fmt"
)

func getString() string {
  return "Test String"
}

func strCat(str1 string, str2 string) string {
  return str1 + str2
}

func Sum(a int, b int) int {
  return a + b
}

func Mul(a int, b int) int {
  return a * b
}

func main() {
  fmt.Println(getString())
  fmt.Println(strCat("strCat ", "TEST"))
  fmt.Println("1+2 = ", Sum(1, 2))
  fmt.Println("3*4 = ", Mul(3, 4))
}
