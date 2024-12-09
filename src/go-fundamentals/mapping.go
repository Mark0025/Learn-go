package main

import "fmt"

func main() {
    // Map in Go
    myMap := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }

    fmt.Println(myMap["Alice"]) // Access value by key
    myMap["Alice"] = 31         // Update value
    fmt.Println(myMap)
}

