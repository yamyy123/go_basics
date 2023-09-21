package main

import "fmt"
type School struct{
       name string 
	   class string
	   fees_status string
	   Discipline Discipline
}

type Discipline struct{
  behaviour string
  respect string
}

func (e School)empinfo()(){
        fmt.Println(e.name)
}
//since we cannot able to add method in structures so we are use a method with a receiver type of the structure in a method outside structure
func main(){
   
	yameen:= School{"yameen","A","paid",Discipline{"fair", "yes"}}
		fmt.Println(yameen)
		yameen.empinfo()
}