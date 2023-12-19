package main

import (
	"fmt"
)

//  Variable Declaration in golang


func main(){

	///////////////////////////////////////////////////////////////

	var iNo_1 int = 10
	var iNo_2 = 12
	iNo_3 := 20
									///////// Three ways in which we can declare variable in Go
	fmt.Println(iNo_1)
	fmt.Println(iNo_2)
	fmt.Println(iNo_3)


	///////////////////////////////////////////////////////////////
	// Data types   
	/*
		int            8/16/32/64    or uint : 8/16/32/64
		float          32/64
		string
		byte
		bool
	*/

	var Number int = 10
	var String string = "Demo"
	var Float float32 = 12.36
	var Byte byte = 'D'
	var Boolean bool = true

	fmt.Println(Number)
	fmt.Println(String)
	fmt.Println(Float)
	fmt.Println(Byte)
	fmt.Println(Boolean)

	/////////////////////////////////////////////
	/// varibale decalare but not initialise

	var Value int

	println(Value)

	///////////////////////////////////////////
	//value  assignmnet after declaration

	var Demo string

	Demo = "Devendra"
	fmt.Println(Demo)

	//////////////////////////////////////////
	// Declare Multiple variables

	var a , b, c, d int = 1,2,3,4
	fmt.Println(a,b,c,d)

	/////////////////////////////////////////
	// Declare Multiple variables but type not specified

	var A,B = 12.4,"Demo"
	println(A,B)

	/////////////////////////////////////////
	// Declare in var block

	var(
		value_1 int
		value_2 string = "Demo"
		value_3 int = 12
	)

	println(value_1,value_2,value_3)
	

	////////////////////////////////////////
	// Namming Rule

	/*
		Camel Case : Each word, except the first, starts with a capital letter:       myVaribleName := "Demo"


		Pascal Case : Each word starts with a capital letter:						  MyVariableName := "Demo"


		Snake Case : Each word is separated by an underscore character:				  my_variable_name := "Demo"


	*/	

	/////////////////////////////////////////
	// Costant     ------ Unchangeable and Read-only
	/*
		There Are two types of Constant as
		1. Typed constant
		2. Untyped Constant
	*/

	const value = 10;             // untyped Const
	println(value)

	const X string = "Demo"      // typed Const
	println(X)

	//////////////////////////////////////////////////
	// Multiple Constant Declaration

	const(
		x int = 10
		y string = "Demo"
		z float32 = 12.23
	)

	println(x,y,z)

	/////////////////////////////////////////
	// Output 
	/*
		Three Options in output
		1. Print()
		2. Println()
		3. Printf()
	*/

	var m,n,q = 10,20,"demo"

	fmt.Print(m)
	fmt.Print("\n")
	fmt.Println(n)
	fmt.Printf(q)

	fmt.Println()

	//////////////////////////////////////
	// Formating verbs
	/*
		%v	Prints the value in the default format
		%#v	Prints the value in Go-syntax format
		%T	Prints the type of the value
		%%	Prints the % sign
	*/

	var M ,N, Q,R = "demo",1,12.23,"%"

	fmt.Printf("%%")
	fmt.Printf("%v",M)
	fmt.Printf("%#v",M)
	fmt.Printf("%T",M,N,Q,R)
	fmt.Println()

	///////////////////////////////////////////////
	// Array

	var Arr = [4] int {10,20,30,40}                 //array length defined

	fmt.Println(Arr)

	var Brr = [...] int {10,20,30,40,50}          // array length inferred

	fmt.Println(Brr)

	Crr := [3] int {10,20,30}					// := sign

	fmt.Println(Crr)

	Drr := [...] string {"BMW","XUV","VOLVO","FORD"}

	fmt.Println(Drr)

	println(Arr[0])

	/////////////////////////////////////////////////
	// Array Initialised

	var Demo_Arr = [4] int {}                       // not initialised
	var Demo_Brr = [4] int {10,20}					// partially initialised
	var Demo_Crr = [4] int {10,20,30,40}			// fully initialised
	var Demo_Drr = [4] int {1:45,3:48}              // Initialised Only specific Elements

	fmt.Println(Demo_Arr,Demo_Brr,Demo_Crr,Demo_Drr)

	/////////////////////////////////////////////////
	// Array Length

	var Demo_arr = [4] int {10,20,30,10}
	fmt.Println("length of array: ",len(Demo_arr))



	//////////////////////////////////////////////////
	///   Slice
	/*
		slice are similar to array but more powerful and flexiable

		three ways foe slice creation:
			1.slice_name := []datatype{values}  

			2.Create a Slice From an Array                 var myarray = [length]datatype{values} 	// An array
														   myslice := myarray[start:end] 			// A slice made from the array

			3.Create a Slice With The make() Function        slice_name := make([]type, length, capacity)

	*/	
	//////////////////////////////////////////////////////////
	// Slice using 			1.slice_name := []datatype{values}  

	slice_arr := [] int {10,20,30,40}
	fmt.Println(slice_arr)
	fmt.Println(len(slice_arr))
	fmt.Println(cap(slice_arr))

	//////////////////////////////////////////////////////////
	//Create a Slice From an Array 

	Demo_Zrr := [7] int {10,20,30,40,50,60,70}
	fmt.Println(Demo_Zrr)

	slice_zrr := Demo_Zrr[2:5]

	fmt.Println(slice_zrr)
	fmt.Println(len(slice_zrr))
	fmt.Println(cap(slice_zrr))


	////////////////////////////////////////////////////////
	/// slice uisng make function

	Demo_slice := make([]int,4,5)
	fmt.Println(Demo_slice)
	fmt.Println(len(Demo_slice))
	fmt.Println(cap(Demo_slice))

	/////////////////////////////////////////////////////////
	// Access, Change, Append and Copy Slices
	// copy(dest, src)


	Access_Demo := [] int {10,20,30}
	fmt.Println("Access 1 index  : ", Access_Demo[1])

	Change_Demo := []int{10,20,30,40}
	fmt.Println(Change_Demo)

	Change_Demo[1] = 50
	fmt.Println(Change_Demo)

	Append_Demo_1 := []int {10,20,30}
	Append_Demo_1 = append(Append_Demo_1,40,50,60)

	fmt.Println(Append_Demo_1)

	Append_Demo_2 := []int {10,20,30}
	Append_Demo_2 = append(Append_Demo_2,Append_Demo_1...)

	fmt.Println(Append_Demo_2)

	//--------------------------------------------------

	numbers := []int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}

	needed_numbers := numbers[:len(numbers)-10]

	Copy_slice := make([]int,len(needed_numbers))

	copy(Copy_slice,needed_numbers);

	fmt.Println(Copy_slice)

}	