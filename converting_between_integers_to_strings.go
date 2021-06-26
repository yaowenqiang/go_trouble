package main

import (
    "fmt"
    "strconv"
    "regexp"
    "log"
)


func main() {
    num := 65
    //str := string(num)
    str := strconv.Itoa(num)

    fmt.Printf("num: %T %d\n", num, num)
    fmt.Printf("num: %T %s\n", str, str)
    //pat := regexp.MustCompile("\d+")
    //pat := regexp.MustCompile("\\d+")
    pat := regexp.MustCompile(`\d+`)

    fmt.Printf("Match: %s\n", pat.FindString(str))

    //strf := "11.23"
    strf := "-3.5762786885e-07"

    numf , err := strconv.ParseFloat(strf, 32)

    if err != nil {
        log.Fatalf("Failed to parse %q: %s\n", str, err)
    }

    fmt.Printf("%T %.20f\n", numf, numf)
}

