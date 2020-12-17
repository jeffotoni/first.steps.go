package main

import "time"

func timer(d time.Duration, i int) <-chan int {
    c := make(chan int)
    go func() {
        time.Sleep(d)
        c <- i
    }()
    return c
}

func main() {
    for i := 0; i < 24; i++ {
        c := timer(1*time.Second, i)
        println(<-c)
    }
}
