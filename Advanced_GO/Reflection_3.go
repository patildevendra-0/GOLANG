package main

import (
	"fmt"
	"reflect"
)

func main(){

	Data := 21

	fmt.Println("DATA BEFORE CHANGING : ",Data)

	Reflection_obj := reflect.ValueOf(&Data).Elem()      //      &Data means address of actual data and 
														//		.Elem() function access that data from that address
														//      means actual value is pointer pointing to that address  	
	Reflection_obj.SetInt(51)

	fmt.Println("DATA AFTER CHANGING : ",Data)

}