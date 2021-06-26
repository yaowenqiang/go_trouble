package main
import (
    "fmt"
    "log"
    _"os"
    "io"
    "net/http"
    "io/ioutil"
)

//func writeMessage(output *os.File)  {
func writeMessage(output io.Writer)  {
    msg := "the best message you'll ever see!\n"
    _, err := output.Write([]byte(msg))

    if err != nil {
        log.Printf("Couldn't write message: %s", err)
    }

    c, ok := output.(io.Closer)

    if !ok {
        fmt.Println("not a closer: quitting")
        return
    }

    fmt.Println("closing output")
    c.Close()
}


func main() {
    f, err := ioutil.TempFile("", "gopher")
    if err != nil {
        log.Printf("Couldn't get temp file : %s", err)
        return
    }

    writeMessage(f)

    fmt.Printf("Write to %s\n", f.Name())

    http.HandleFunc("/", index)
    fmt.Println("starting server...")
    http.ListenAndServe("localhost:3333", nil)

}


func index(w http.ResponseWriter, r *http.Request) {
//    w.Write([]byte("thanks for visiting!"))
    writeMessage(w)
}
