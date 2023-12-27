package main

import "fmt"

type Engine struct{
	HorsePower int
}

type Wheels struct{
	Size int
}

type Car struct{              // car struct embedding engine struct and wheels struct 
	Engine
	Wheels
	Model string
}

func main(){

	Cobj := Car{
		Engine: Engine{HorsePower: 200},
		Wheels: Wheels{Size: 10},
		Model: "Sedan",
	}

	fmt.Println(Cobj)

}