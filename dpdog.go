package main

import "fmt"
import "time"


func main() {
    fmt.Println("Hello world")
    t1 := time.NewTimer(time.Second * 5)
    t2 := time.NewTimer(time.Second * 10)

    for {
        select {

        case <-t1.C:
            println("5s timer")
            t1.Reset(time.Second * 5)

        case <-t2.C:
            println("10s timer")
            t2.Reset(time.Second * 10)
        }
    }
}
