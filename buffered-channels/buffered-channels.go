package main

import "fmt"

func main() {

        ch := make(chan int, 2)

        ch <- 23
        ch <- 42
        fmt.Println(<-ch)
        fmt.Println(<-ch)
}

