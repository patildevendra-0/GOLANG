package main

import (
	"fmt"
	"sync"
)


func Demo_1(channel chan <-string,WG *sync.WaitGroup){                  // send only channel

	defer WG.Done()

	Data := "SHREE GANESH"

	channel <- Data

}

func Demo_2(channel <-chan string,WG *sync.WaitGroup){                  // Recive only channel
	
	defer WG.Done()

	Data := <-channel

	fmt.Println("DATA RECIVED	: ",Data)

}	



func main(){

	var WG sync.WaitGroup

	Channel := make(chan string)

	WG.Add(2)

	go Demo_1(Channel,&WG)
	go Demo_2(Channel,&WG)

	WG.Wait()

	close(Channel)

}