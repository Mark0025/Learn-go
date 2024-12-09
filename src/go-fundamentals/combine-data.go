package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {Name: "Alice", Age: 30},
        {Name: "Bob", Age: 25},
    }

    for _, person := range people {
        fmt.Printf("%s is %d years old.\n", person.Name, person.Age)
    }
}

