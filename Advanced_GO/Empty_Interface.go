package main

import "fmt"

func Demo(data interface{}){                          // this function handle any type of data
	fmt.Println("DATA FROM DEMO : ",data)
}

func main(){

	Demo(10)
	Demo("demo")
	Demo(12.3)
}