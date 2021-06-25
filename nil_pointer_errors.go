package main

import (
    "net"
    "log"
    "fmt"
)


type evenNum struct {
    val int
}


func (en evenNum) half() int {
    return en.val / 2
}

func getEvenNum(i int) *evenNum {
    if i % 2 == 1 {
        return nil
    }

    return &evenNum{val: i}
}



func main() {

    eNum := getEvenNum(3)

    fmt.Println(eNum.half())

    conn, err := net.Dial("tcp", "localhost:80")
    if err != nil {
        log.Printf("Failed to open connection: %s", err)
        log.Printf("the conn is type: %T", conn)
        return
    }
    conn.Write([]byte("hello"))
}

