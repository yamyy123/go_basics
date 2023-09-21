package main

import "fmt"

func main(){
	increment:=king()
	fmt.Println(increment())
	fmt.Println(increment())
}

func king()func()int{
	count:=0
	return func()int{
	  count++
      return count
	}
}

//in this the anonymous function inside the king function returns the function in to the increment varible whenever the increment is called 
//the previous returned value will be stored and the anonymous func can be called using that assigning identifier

//closure means even after the execution of the outer function we can able to acces the inner func with creating the instance of the function 
//with the previous stored information