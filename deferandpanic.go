package main

import "fmt"

func main(){
	defer fmt.Println("hello")
	defer fmt.Println("bye")//it executes like stack lifo(lastin firstout manner)
	fmt.Println("up of panic")
	panic("die")//panic will suddenly stops the execution of code but the defer statements above the panic will be 
	//executed but downside will not be printed
	fmt.Println("down of panic")
	defer fmt.Println("down of panic defer")
}