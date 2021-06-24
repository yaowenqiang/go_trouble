package main
import (
    "fmt"
    "os"
)

var msg string

func init() {
    msg = os.Getenv("message")
    //export message="hello world"
}

func main() {
    fmt.Println(msg)
}
