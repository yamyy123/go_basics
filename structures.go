package main

import "fmt"


type Address struct{
     name string
	 pincode int64
}

func main(){
	// var saru Address
	// saru.name="hi"
	// saru.pincode=638401
	// var saru Address=Address{"yameen",638401}
	//var saru =Address{"yameen",638401}
	/* both also valid declarations in go*/
	// saru:=Address{"yameen",638401}
	// fmt.Println(saru)
	//creating a instance of structure using struct variable
	// rect1 := new(Address)
    //  rect1.name="yameen"
	//  fmt.Println(rect1)
//(or)

//var rect1 = new(rectangle)
//creating the instance of the struct by the pointer variable
      var yameen=&Address{"yameen",638401}
	  fmt.Println(yameen)
}