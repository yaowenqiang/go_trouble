package main
import (
    "fmt"
    "sync"
)

var count int
var mu sync.Mutex

func incr(wg *sync.WaitGroup)  {
    mu.Lock()
    defer mu.Unlock()
    for i := 0; i < 10000;i++ {
        count++
    }

    printCount("incr")
    wg.Done()
}


func printCount(label string)  {
    fmt.Printf("%s coun : %d\n", label, count)
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 5;i++ {
        wg.Add(1)
        go incr(&wg)
    }

    wg.Wait()

    printCount("final!")

}

// for i in {1..3}; do go run main.go; echo; done
