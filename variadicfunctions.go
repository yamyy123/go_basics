package main

import "fmt"

func variadicfunc(s...string){
	fmt.Println(s[0])
	fmt.Println(s[1])


}
func main(){
	variadicfunc("hi","bye","tommorow","dietoday")
}