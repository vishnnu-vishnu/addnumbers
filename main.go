package main

import (
    "fmt"
    "github.com/vishnnu-vishnu/addnumbers/adder"
)

func main() {
    a := 5
    b := 3
    sum := adder.Add(a, b)
    fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)
}
