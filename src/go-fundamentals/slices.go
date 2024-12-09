package main

import "fmt"

func main() {
    // Create a slice (dynamic array)
    slice := []int{1, 2, 3}
    slice = append(slice, 4)        // Append a single value
    slice = append(slice, 5, 6)     // Append multiple values
    fmt.Println(slice)
}

