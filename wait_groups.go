package main
import (
    "fmt"
    "sync"
    "github.com/shawnmilo/girya"
)

func msg(i int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println(i)
    kg := 20.0
    fmt.Printf("%.1f kg is %.1f lbs.\n", kg, girya.KgToLbs(kg))
}
func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(num int) {
            msg(num, &wg)
        }(i)
    }

    wg.Wait()
    fmt.Println("Done")
}

