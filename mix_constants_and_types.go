package main

import (
    "fmt"
    "time"
)

const a = 1
var b = 2
var c int = 3
var e int32 = 5


func main() {
    d := 4
    var du  time.Duration = 2
    fmt.Printf("%T %T %T %T\n", a, b, c, d)

    for i := 1; i < 2; i++ {
        fmt.Printf("%d + %d = %d\n", i, a, i+a)
        fmt.Printf("%d + %d = %d\n", i, b, i+b)
        fmt.Printf("%d + %d = %d\n", i, c, i+c)
        fmt.Printf("%d + %d = %d\n", i, d, i+d)
        //fmt.Printf("%d + %d = %d\n", i, e, i+e)
        fmt.Printf("%d + %d = %d\n", i, e, i+int(e))
    }

    dur := time.Duration(b)
    //time.Sleep(a)
    //time.Sleep(b)
    time.Sleep(dur)
    time.Sleep(du)
    //time.Sleep(int64(b))
    fmt.Println("done")
}

