package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")  // prints a line with a newline character at the end
    fmt.Print("Hello, ")  // prints without a newline character at the end
    fmt.Print("World!\n")  // you can add a newline character manually if needed

    name := "oki"
    age := 17
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)  // formatted printing
}
