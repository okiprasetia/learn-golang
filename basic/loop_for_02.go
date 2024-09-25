package main

import "fmt"

func main() {
  sum := 1
  for i := 1; i <5; i++ {
    sum += 1
    fmt.Println(sum)
  }
}