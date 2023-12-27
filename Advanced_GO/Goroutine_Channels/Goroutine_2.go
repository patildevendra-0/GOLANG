package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintNumber(start int ,end int, WG *sync.WaitGroup){

	defer WG.Done()
	for i := start; i <= end; i++ {
		fmt.Println(i)
		time.Sleep(2*time.Second)                       /// this not essential but to watch how rotine schedule one by one for that ..... see in output..
	}
	fmt.Println("---------------------------------")
}

func main(){
 
	var WG sync.WaitGroup

	WG.Add(2)

	go PrintNumber(1,10,&WG)
	go PrintNumber(11,20,&WG)
	
	WG.Wait()

	fmt.Println("END OF MAIN...")
}