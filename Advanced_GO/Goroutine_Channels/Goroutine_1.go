package main

import (
	"fmt"
	"sync"
)

func PrintNumbers(WG *sync.WaitGroup){

	defer WG.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func PrintCharachter(WG *sync.WaitGroup){

	defer WG.Done()

	for char := 'A'; char <= 'Z'; char++ {
		fmt.Printf("%c\n",char)
	}

}



func main(){

	var WG sync.WaitGroup

	WG.Add(2)

	go PrintNumbers(&WG)
	go PrintCharachter(&WG)

	WG.Wait()
	
	fmt.Println("MAIN FUNCTION END")

}