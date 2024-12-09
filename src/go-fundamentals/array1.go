package main

import "fmt"

func main() {
    // Fixed-size array in Go
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr[0]) // Access the first element

    arr[0] = 10         // Modify the first element
    fmt.Println(arr)
}

