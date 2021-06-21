package main
import (
    "fmt"
)

func getNumber() int {
    fmt.Printf("The number is %d\n", num)
    num := 42
}

func main() {
    num := getNumber()
    fmt.Println(num)
}

