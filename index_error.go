package main
import (
    "fmt"
)

func main() {
    things := []string{"screwdriver", "torch", "phone booth"}

    fmt.Println(things)
    fmt.Printf("%#v\n", things)
    fmt.Printf("Item 0 : %s\n", things[0])
    fmt.Printf("Item 2 : %s\n", things[2])

    for i, item := range things {
        fmt.Printf("Item %d : %s\n", i, item)
    }

    for i := 0; i < 5; i++ {
        if i > len(things) -1 {
            fmt.Printf("No index %d\n", i)
            continue
        }
        fmt.Printf("Item %d : %s\n", i, things[i])
    }


}

