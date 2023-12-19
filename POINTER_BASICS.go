package main

import "fmt"

func Pointer_Function(M *int){

	*M = 123
}

func address_function(M *int){
	*M = 456
}

type Employee struct{
	name string
	empid int
}

func main(){

	var iNo int = 10

	var P *int = & iNo

	fmt.Println("value of iNo",iNo)
	fmt.Println("value of p",P)
	fmt.Println("adress of iNo" ,&iNo)
	fmt.Println("value of * p : ",*P)


	*P = 123

	var Q **int = & P

	fmt.Println("value of iNo after change: ",iNo)
	fmt.Println("value of **Q" ,**Q)

	/////////////////////////////////////////////////////////////////////////////////////
	////  passing pointer to function
	var Demo = 10
	var N *int = &Demo
	fmt.Println("value of n before function call  :",*N)

	Pointer_Function(N)

	fmt.Println("value of n after function call  :",*N)

	/////////////////////////////////////////////////////////////////////////////////////////
	//// passing address to function

	iNo_1 := 10
	fmt.Println("Before call: ",iNo_1)
	address_function(&iNo_1)
	fmt.Println("after call: ",iNo_1)

	/////////////////////////////////////////////////////////////////////////////////////////
	/// pointer to structure

	var eobj Employee 
	eobj.name = "xyz"
	eobj.empid = 123

	Z := &eobj

	fmt.Println("Z pointer point to struct : ",Z);

	Z.name = "12321"
	fmt.Println("Z pointer point to struct : ",Z);

	//////////////////////////////////////////////////////////////////////////////////////////////
	/// compare two pointers == or !=

	ino_1 := 10

	Ptr := & ino_1
	Qtr := & ino_1
	Ztr := & iNo

	bret := Ptr == Qtr
	fmt.Println(bret)

	bRet := Ptr != Ztr
	fmt.Println(bRet)

	cRet := Ptr == Ztr
	fmt.Println(cRet)





}