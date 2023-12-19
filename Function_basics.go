package main

import "fmt"

func Demo_Function(){
	fmt.Println("SHREE GANESH ...")
	fmt.Println("INSIDE DEMO FUNCTION...")
}

func Single_parameter(ino int){
	ino++
	fmt.Println("value increment : ",ino)
}

func Multiple_Parameter(ino_1 int,ino_2 int){
	var ans int = 0;
	ans = ino_1+ino_2
	fmt.Println("Addition from helper function :",ans)
}

func Return_Function(ino_1 int,ino_2 int) int{
	var ans int = 0
	ans = ino_1+ino_2
	return ans
}

func Multiple_Return(ino_1 int,ino_2 int)(int, int){
	Ans_1:= ino_1+ino_2
	Ans_2:= ino_1-ino_2

	return Ans_1,Ans_2
}

func main(){

	///////////////////////////////////////////
	// Function with No parameter
	Demo_Function()

	/////////////////////////////////////////////
	//function with single parameter
	Single_parameter(10)

	/////////////////////////////////////////////
	// Multiple parameter
	Multiple_Parameter(10,20)

	/////////////////////////////////////////////
	// Function with retunn value
	iRet := Return_Function(10,20)
	fmt.Println("Addition from main: ",iRet)

	//////////////////////////////////////////////
	// Function With multiple return value

	iRet_1 , iRet_2 := Multiple_Return(10,20);
	fmt.Println("Addition and Substration from main : ", iRet_1,iRet_2)

	///////////////////////////////////////////////
	//



}