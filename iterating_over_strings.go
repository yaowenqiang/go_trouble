package main

import (
    "fmt"
)

func main() {
    world := "世界"
    世界 := "world"
    fmt.Printf("hello, %s\n", 世界)
    fmt.Printf("hello, %s\n", world)
    describe(world)
    describe(世界)

    for i, val := range world {
        fmt.Printf("%d %s\n", i, string(val))
    }
}


func describe(s string) {
    fmt.Printf("String %q, length %d\n", s, len(s))

    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
        fmt.Printf("Char: %d, %s\n", s[i], string(s[i]))
    }

    for _, r := range []rune(s) {
        //fmt.Println(s[i])
        fmt.Printf("Rune: %d, %s\n", r, string(r))
    }

    junk  := "你在哪里"

    fix(junk)
}



func fix(s string) {
    var fixed []byte

    for _, c := range s {
        fixed = append(fixed, byte(c))
    }

    fmt.Printf("Fixed: %s\n", string(fixed))
}


