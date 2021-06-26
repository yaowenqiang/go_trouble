package main
import (
    "fmt"
    "time"
    "log"
)

func main() {
    fmt.Println(time.Now().Format("3:04:05"))
    fmt.Println(time.Now().Format("15:04:05"))

    d1, err := time.Parse("2006-01-02", "1955-11-05")
    if err != nil {
        log.Fatalf("Sorry, Doc: %s\n", err)
    }

    d2, err := time.Parse("2006-01-02", "1985-10-21")
    if err != nil {
        log.Fatalf("Sorry, Marty: %s\n", err)
    }

    fmt.Printf("To %s\n", d1)
    fmt.Printf("Back To %s\n", d2)


    fmt.Println(d1.Format(time.RFC1123Z))
    fmt.Println(d2.Format("January 2, 2006"))

    fmt.Println(d2.Sub(d1))
    fmt.Println(d2.Add(time.Hour * 24))
    fmt.Println(d2.Add(-time.Hour * 24))


}

