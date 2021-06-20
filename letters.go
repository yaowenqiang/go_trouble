
package main

import (
    "fmt"
)
func main() {
    for _, c := range []rune("gopher!") {
        fmt.Println(string(c))
    }
}
