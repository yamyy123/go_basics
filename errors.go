package main

import (
	"errors"
	"fmt"
)
//err has bulitin interface as

/*
type error interface {

Error() string

} it will either gives err or nil
*/
func main(){
	
    err := errors.New("math - square root of negative number")
	fmt.Println(err)
	fmt.Println("hi")
}