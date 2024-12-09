package main

import "fmt"

func main() {
    // For loop (Go has no while loop)
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Equivalent to while loop in Python
    count := 0
    for count < 5 {
        fmt.Println(count)
        count++
    }
}

