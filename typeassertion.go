package main

import (
    "fmt"
)

func assert(s interface{}) { //type asse
   // i:=s.(int)//if we write like this panic occurs
//    i:= s.(string)
// 	fmt.Println(i)//this is correct

	//to avoid panic this can be used

	v, ok := s.(int)
    fmt.Println(v, ok)
	//ok will give true or false and v value will be zero 
	//this is the solution to prevent panic
}

func main() {
   var s interface{}="string"
   assert(s)
}
