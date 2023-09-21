package main

import (
	"errors"
	"fmt"
)

func main(){
	value,err:=squareroot(0)
	if err!=nil{
		fmt.Println("error has been occured:",err)
	}else{
		fmt.Println(value)
	}
}

func squareroot(a int)(int,error){
	if a<=0{
		return 0,errors.New("byebye")
	}else{
		return 1,nil
	}
}
