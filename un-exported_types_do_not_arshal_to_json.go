package main
import (
    "fmt"
    "log"
    "encoding/json"
)

type person struct {
    Name string
    Phone string
}

func main() {
    p := person{Name:"Jack", Phone: "12345"}
    j, err := json.Marshal(p)

    if err != nil {
        log.Fatalf("Failed to marshal p: %s\n", err)
    }

    fmt.Printf("Original: %+v\n", p)
    fmt.Printf("JSON output: %s\n", string(j))
}

