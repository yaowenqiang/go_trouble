package main

import (
    "fmt"
)

func main() {
    primes := []int{2,3,7,11,13,17}
    for i,p := range primes {
        fmt.Printf("item %d is %d\n", i,p)
    }
}
