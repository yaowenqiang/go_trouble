package main
import (
    "fmt"
    "net"
    "time"
)
func main() {
    msg := "Message 42: Don't panic!"
   // conn, _ := net.DialTimeout("tcp", "localhost:22", time.Second)
    //conn.Write([]byte(msg))
    conn, err := net.DialTimeout("tcp", "localhost:22", time.Second)
    if err != nil {
        fmt.Println("Failed to create connection: %s\n", err)
        return
    }
    conn.Write([]byte(msg))
    fmt.Println("All Done!")
}

