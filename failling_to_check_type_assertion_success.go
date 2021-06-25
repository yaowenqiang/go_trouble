package main

import (
    "fmt"
)

type animal struct {
    name string
    sound string
}

func (a animal) makeNoise() {
    fmt.Printf("the %s goes %q.\n", a.name, a.sound)
}


func print(thing interface{}) {
    fmt.Printf("thing is a %T with a value of %#v\n", thing, thing)
    //thing.makeNoise()
    x, ok := thing.(animal)
    if ok {
        fmt.Printf("x is %T with a value of %#v\n", x, x)
        x.makeNoise()
    } else {
        fmt.Printf("thing %v is not an animal\n", thing)
    }

}

func main() {
    a := animal{
        name: "dog", sound: "bark",
    }
    a.makeNoise()
    print(a)
    print("cat")
}

