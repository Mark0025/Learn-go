package main

import "fmt"

// Struct definition
type Person struct {
    Name string
    Age  int
}

func main() {
    alice := Person{Name: "Alice", Age: 30}
    fmt.Println(alice.Name)
}

