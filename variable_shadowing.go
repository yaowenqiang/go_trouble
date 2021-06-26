package main

import (
    "fmt"
)

type shirt struct {
    size int
}


func main() {
    i := 23
    fmt.Printf("i is %d\n",i)

    colors := []string{"purple", "gray", "green", "black"}
    //for _, color := range colors {
    for i, color := range colors {
        fmt.Printf("Color %d is %s\n", i, color)
        j := 5 * i
        fmt.Printf("j: %d\n",j)
    }


    fmt.Printf("i is %d\n",i)
    //fmt.Printf("j is %d\n",j)

    x := 3
    y := 4

    fmt.Printf("before x : %d, y: %d\n", x, y)
    x, a := 5000, 7000
    //x, y := 5000, 7000, error
    fmt.Printf("x : %d, a: %d\n", x, a)
    fmt.Printf("after x : %d, y: %d\n", x, y)



    s := shirt{}
    var b int
    s.size, a = 5 ,3
    //can't use := on struct at all

    fmt.Printf("size: %d, b : %d\n", s.size, b)



}
