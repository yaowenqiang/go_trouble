package main

import (
    "fmt"
    "time"
)

func main() {
    m := make(map[int]int)

    for i := 0; i < 100; i++ {
        go func(i int) {
            m[i] = i
        }(i)
    }

    time.Sleep(time.Second)
    fmt.Printf("map: %+v\n", m)
}
