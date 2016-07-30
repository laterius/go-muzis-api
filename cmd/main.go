package main

import (
	"fmt"
	_ "github.com/laterius/go-muzis-api"
)

type MyStruct struct {
	A int
}

func SetA(myS MyStruct) MyStruct {
	myS.A = 10
	return myS
}

func SetA2(myS *MyStruct) {
	myS.A = 10
}

func main(){
	//arr1 := [3]int {1, 2, 3}
	slice1 := []int{}

	slice2 := slice1

	slice1[0] = 11
	fmt.Println(slice2[0])
}
