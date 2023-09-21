package main

import (
	"fmt"
	"sync"
	//"time"
)

var counter = 0
var wg sync.WaitGroup

func increment() {
    for i := 0; i < 1000; i++ {
        counter++
    }
    wg.Done()
 
}

func main() {
    wg.Add(5)
    go increment()
    go increment()
	go increment()
	go increment()
	go increment()
   wg.Wait()
   //time.Sleep(1*time.Second)
    fmt.Printf("Counter: %d\n", counter)
}