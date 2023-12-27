package main

import (
	"fmt"
)

type Printer interface{
	Print()
}

type Saver interface{
	Save()
}

type Document interface{                // Document interface embedding printer and saver interface
	Printer 
	Saver
}

type MyDocument struct{
	content string
}

func (MD MyDocument) Print(){                                  

	fmt.Println("Data Printing  : ",MD.content)

}

func (MD MyDocument) Save(){

	fmt.Println("Data Saving : ",MD.content)

}

func main(){

	myDoc := MyDocument{content: "hello golang"}

	myDoc.Print()
	myDoc.Save()	

}