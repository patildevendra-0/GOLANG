package main

import (
	"fmt"
	"sync"
)


func Producer(Data string,channel chan string, WG *sync.WaitGroup){

	defer WG.Done()
	
	channel <- Data
	fmt.Println("-----------------INSIDE PRODUCER----------------------------------")

	Message := <-channel
	fmt.Println("DATA RECIVED FORM RECIVER : ",Message)

}

func Reciver(channel chan string,WG *sync.WaitGroup){
	defer WG.Done()

	fmt.Println("---------------------INSIDE RECEIVER------------------------------")
	Data := <-channel

	fmt.Println("DATA FROM PRODUCER : ",Data)

	channel<-"OK BYE BYE"
}


func main(){

	var WG sync.WaitGroup

	Channel := make(chan string)

	Data := "SHREE GANESH"

	WG.Add(2)

	go Producer(Data,Channel,&WG)

	go Reciver(Channel,&WG)

	WG.Wait()

	close(Channel)

}