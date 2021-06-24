package main
import (
    "fmt"
    "io/ioutil"
)


func main() {
    msg := "Go can read and write files\n"

    filename := "/etc/messages.txt"
    //ioutil.WriteFile(filename, []byte(msg), 5644)
    err := ioutil.WriteFile(filename, []byte(msg), 5644)
    if err != nil {
        fmt.Println("Failed to write file: %s\n", err)
    }

    fmt.Println("All done")
}

