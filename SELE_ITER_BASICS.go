package main

import "fmt"

/////////////////////////////////////
// In Golang only one loop is present which is for loop

func FOR_LOOP(){
	var iCnt int = 0;
	
	for iCnt =1; iCnt <= 10; iCnt++ {

		fmt.Println(iCnt)	
	}
}

//////////////////////////////////////
// if else

func Selection_Demo(ino int)bool{

	if ino % 2 == 0{
		return true
	}else{
		return false
	}
	
}

//////////////////////////////////////
// switch case

func Switch_demo(ino int){

	switch ino {
	case 1: fmt.Println("---------------1")
	case 2: fmt.Println("---------------2")
		
	}
	
}


func main(){

	FOR_LOOP()
	bRet := Selection_Demo(10)
	if bRet == true{
		fmt.Println("even number")
	}else
	{
		fmt.Println("odd number")
	}

	Switch_demo(1)
}