package main

import (
	"fmt"
	"time"
)

// func main() {
// 	defer func() {
//        if r:=recover();r!=nil{
// 		fmt.Println("jilapi",r)
// 	   }
// 	}()
// 	panic("err")
// }

func main() {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered from goroutine panic:", r)
            }
        }()
        // Simulate a panic in a goroutine
        panic("Panic in goroutine")
    }()

    time.Sleep(1 * time.Second) // Give the goroutine some time to run

    fmt.Println("Main function completed.")
}



