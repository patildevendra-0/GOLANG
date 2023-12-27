package main

import (
	"fmt"
)

func AssertionDemo(Data interface{}){             // this function identify which type of data comes in interfcae 
												 // or which type of data is given to funcion while calling
	if Ret , Ok := Data.(int) ; Ok{
		fmt.Println("THIS IS INTEGER : ",Ret)
	}else if Ret,Ok:=Data.(string);Ok {
		fmt.Println("THIS IS STRING : ",Ret)
	}else if Ret,Ok:=Data.(float64);Ok{
		fmt.Println("THIS IS FLOAT :",Ret)
	}
		


}

func main(){

	AssertionDemo(10)
	AssertionDemo("Shree Ganesh")
	AssertionDemo(12.05)

}