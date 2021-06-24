package main
import (
    "fmt"
    "log"
    "io/ioutil"

)

func main() {
    data, err := ioutil.ReadFile("db.go")
    if err != nil {
        log.Fatalf("Failed to read file: %s\n", err)
    }

    //chars := make(map[rune]int)
    chars := make(map[rune]struct{})

    for _, r := range([]rune(string(data))) {
        //chars[r]++
        chars[r] = struct{}{}
    }

    /*
    for r, c := range chars {
        fmt.Printf("%s: %d\n", string(r), c)
    }

    */

    for r, _ := range chars {
        fmt.Printf("%s ", string(r))
    }
    //fmt.Println(string(data))
}

