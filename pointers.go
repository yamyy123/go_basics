package main

import "fmt"
//The default value or zero-value of a pointer is always nil.
func main(){
	var i int =0
	var a *int=&i
	fmt.Println(*a)
	fmt.Println(a)
	var y *int
	fmt.Println(y)//the value of this is nil

	u:=48
	o:=&u
	fmt.Println(*o)
	var k string="pol"
	var l *string=&k
	fmt.Println(*l)
}