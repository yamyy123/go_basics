package main

import "fmt"


func types(i interface{}){
	switch x:=i.(type){
	case int:
		fmt.Printf("integer is %d",x)
	case string:
	    fmt.Printf("string is %s",x)
	default:
		fmt.Println("unknown type")
	}

}
func main(){
	var i interface{}=56
	var s interface{}="string"
	types(i)
	types(s)
}