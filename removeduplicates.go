package main

import "fmt"

func main(){
	var arr =[]int{1,3,3,3,2,4}
	arr1:=removeduplicates(arr)
	fmt.Println(arr1)
}

func removeduplicates(arr []int)([]int){

	slice:=[]int{}
	mymap:=make(map[int]bool)
	for _,value:=range arr{
		if found:=mymap[value];!found{
			mymap[value]=true
			slice=append(slice,value)
			continue
		}
		element:=mymap[value]
		fmt.Println(element)
	}
	return slice
}