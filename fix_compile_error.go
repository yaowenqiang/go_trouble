package main

import (
    "fmt"
    "log"
    "strconv"
)

//var size 5
var size  = 5

var colors = []string{
    "red",
    "orange",
    "yellow",
    "green",
    "indigo",
    //"viblet
    //"viblet" }
    "viblet",
}

func main() {
    fmt.Printf("the size is %d\n", size)

    for _, c :=  range colors {
        fmt.Println(c)
    }

    s := "32"
    i, err := strconv.Atoi(s)
    if err != nil {
        log.Fatalf("Failed to convert %q to number: %s\n", s, err)
    }
    fmt.Printf("The number is %d\n", i)

    s = "a32"
    i, err = strconv.Atoi(s)
    if err != nil {
        log.Fatalf("Failed to convert %q to number: %s\n", s, err)
    }
    fmt.Printf("The number is %d\n", i)

}
