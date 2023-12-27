package main

import (
	"fmt"
	"reflect"
)


func main(){

	sValue := "Shree Ganesh"
	iValue := 11
	fValue := 12.45

	Reflection_Type := reflect.TypeOf(sValue)
	fmt.Println("TYPE OF sValue  : ",Reflection_Type)

	Reflection_Type = reflect.TypeOf(iValue)
	fmt.Println("TYPE OF iValue  : ",Reflection_Type)

	Reflection_Type = reflect.TypeOf(fValue)
	fmt.Println("TYPE OF fValue  : ",Reflection_Type)
}