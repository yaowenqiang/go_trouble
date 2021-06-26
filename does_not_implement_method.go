package main

import (
    "fmt"
    "log"
    "os"
)

type nums struct {
    a int
    b int
}

type adder interface {
    add()
}

type words struct {
    first string
    second string
}

func (w words) add() {
    fmt.Printf("%s  %s\n", w.first, w.second)
}


func (n nums) add() {
    fmt.Printf("%d + %d = %d\n", n.a, n.b, n.a + n.b)
}

//func worker(n nums) {
func worker(n adder) {
    n.add()
}

func main() {
    m := nums{a: 1, b: 2}
    m.add()
    worker(m)

    w := words{first: "maximum", second: "effort"}
    //log.SetOutput(w)
    log.SetOutput(os.Stdout)
    //w.concat()
    log.Println("done")

    worker(w)
}
