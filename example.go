package main

import "fmt"

//	func increment()(func()int){
//		number :=1
//		fmt.Println(number)
//		return func()(int){
//			number*=100
//			return number
//		}
//	}
func main() {
	// instance :=increment()
	// fmt.Println(instance())
	// fmt.Println(instance())

      mymap:=make(map[string]int)
	  mymap["cricket"]=10
	  mymap["andril"]=0
	  mymap["zakeer"]=11
	  fmt.Println(mymap)
}
