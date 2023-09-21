package main

import "fmt"

func main(){
	r := 0
func(r int) {
    r++
    fmt.Println(r)
}(r)
}