package main

import (
    "fmt"
)

func main() {
    var x, y float64

    fmt.Print("Enter first number: ")
    fmt.Scan(&x)

    fmt.Print("Enter second number: ")
    fmt.Scan(&y)

    fmt.Printf("Addition: %.2f\n", x+y)
    fmt.Printf("Subtraction: %.2f\n", x-y)
    fmt.Printf("Multiplication: %.2f\n", x*y)
    fmt.Printf("Division: %.2f\n", x/y)
}

