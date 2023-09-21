package main

import "fmt" 

type Schools struct{
	name string 
	class string
	fees_status string

}

type Employee struct{
	name string
	companyname string
}


func (e Schools)empinfo()(){
	 fmt.Println(e.name)
}
/*here for same method name we cannot able to achieve polymorphism like calling different parameters 
we can able to access only with accessing different structures */

func (e Employee)empinfo()(){
	fmt.Println(e.name)
}




//since we cannot able to add method in structures so we are use a method with a receiver type of the structure in a method outside structure
func main(){

 student:= Schools{"yameen","A","paid"}
	 student.empinfo()
  worker:=Employee{"siva","netxd"}
  worker.empinfo()

}