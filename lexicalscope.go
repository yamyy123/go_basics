package main

import "fmt"

func outer() {
	x := 10

	inner := func() {
		x++
		fmt.Println(x) // Accesses the x variable from the outer scope
	}

	inner()
	inner()
}
//here we are accessing the x from the outer to inner function
func main() {
	outer()
	outer()
}