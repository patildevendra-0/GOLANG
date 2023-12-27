package main

import (
	"fmt"
	"reflect"
)

func main(){

	Data := 21

	Value := reflect.ValueOf(Data)      // Value : this give reflection object not actual value output is 21 but this is reflection object

	Actual_Value := Value.Interface()   //  .Interface() give actual value

	fmt.Println("ACTUAL DATA : ",Actual_Value)

}