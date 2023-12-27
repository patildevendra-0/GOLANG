package main

import (
	"fmt"
	"reflect"
)


type Car_demo struct{

	Name  string
	HorsePower int
	Model string
}

func main(){

	Cobj := Car_demo{Name: "BMW",HorsePower: 300,Model: "sidan"}

	fmt.Println("DATA BEFORE CHNAGE : ",Cobj)

	Data := reflect.ValueOf(&Cobj).Elem()

	Data.Field(0).SetString("Mustag")
	Data.Field(1).SetInt(350)

	fmt.Println("DATA AFTER CHANGE : ",Data)



}