package main

import "fmt"

type profile struct{
	firstname string
	lastname string
	phone string 
	pincode string
}

func main(){
	var m=map[string]profile{
       "profile 1":profile{
		"yameen","mohammed","8754851351","638401"},
		"profile 2":profile{
			"sujith","murugan","8754851351","638402"},
	}
	fmt.Println(m)
}