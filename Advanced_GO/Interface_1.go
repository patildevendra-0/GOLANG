package main

import "fmt"

type Shape interface{
	Area() float64
}

type Circle struct{
	Radius float64
}

type Rectangle struct{
	Width float64
	Height float64
}

func (C Circle) Area(){

	Ans := C.Radius * C.Radius *3.14
	fmt.Println("AREA OF CIRCLE IS : ",Ans)
	
}

func(R Rectangle) Area(){

	Ans := R.Height * R.Width
	fmt.Println("AREA OF RECTANGLE : ",Ans)

}


func main(){

	Circle := Circle{Radius: 4}
	Rectangle:= Rectangle{Height: 4,Width: 4}

	Circle.Area()
	Rectangle.Area()

}