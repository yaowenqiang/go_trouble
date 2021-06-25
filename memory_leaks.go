package main
import (
    "time"
    "fmt"
)

func getOne(c chan int64) {
    for {
        c <- 1
    }
}

func main() {
    for {
        t := time.NewTicker(time.Nanosecond)
        <-t.C
        //always remember to close the channel
        t.Stop()
    }
    var total int64
    for {
        c := make(chan int64)
        go getOne(c)

        total += <-c

        if total % 1000000 == 0 {
            fmt.Println(total)
        }
    }
}
