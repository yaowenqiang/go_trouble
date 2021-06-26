package main

import (
    "fmt"
)


type counter struct {
    count int
}

type incrementer interface {
    incr()
}

func (c *counter) incr() {
    //fmt.Printf("befor incr: %d\n", c.count)
    c.count++
    //fmt.Printf("after incr: %d\n", c.count)
}

func main() {
    //c := counter{}
    c := &counter{}
    fmt.Printf("Counter befor: %d\n", c.count)
    //c.incr()
    //tick(c)
    //tick(&c)
    tick(c)
    fmt.Printf("Counter after: %d\n", c.count)
}


func tick(i incrementer) {
    i.incr()
}
