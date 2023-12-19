package main

import "fmt"

type student struct
{
	name string
	rollno int
	age int
	division string
}

func map_demo(){
	value := map[string]int{"bmw":1,"ford":2,"volvo":4}
	fmt.Println(value)
}

func map_demo_2(){

	value := make(map[string]string)

	value["model"] = "bmw"
	value["year"] = "1995"

	fmt.Println(value)
}

func main(){

	var sobj student

	sobj.name = "Devendra"
	sobj.age = 15
	sobj.rollno = 46
	sobj.division = "G"

	fmt.Println(sobj)
	map_demo()
	map_demo_2()

}	