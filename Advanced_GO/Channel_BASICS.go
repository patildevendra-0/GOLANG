package main

import "fmt"

func main(){

	fmt.Println("SHREE GANESH...")

}





/*
	--------------------------------------------------------------------------------------------------------------------
													Channels In Golang
	Sending Data to a Channel (ch <- data):

		ch		: This is the channel to which you're sending data.
		<-		: The arrow indicates the direction of data flow (from the right to the left).
		data	: This is the actual data you're sending into the channel.


		# Types of Channel :
							1.Bidirectional channel ----------------- ch := make(chan int)

							2.Send-Only channel---------------------- ch := make(chan<- string)

							3.Receive-Only Channel------------------- ch := make(<-chan float64)



	--------------------------------------------------------------------------------------------------------------------

	Real-World Analogy  : Restaurant Kitchen

			Imagine you're managing a busy restaurant kitchen with different chefs responsible for various tasks. 
			The kitchen represents your concurrent program, and each chef represents a goroutine. 
			Now, let's introduce channels to see how they facilitate communication.

		Task Parallelization:

			In a restaurant kitchen, chefs can work on different tasks simultaneously. 
			For example, one chef may be chopping vegetables, another may be grilling meat, and a third may be preparing sauces.
			
			// Goroutines (chefs) working concurrently

						go chopVegetables(ch)
						go grillMeat(ch)
						go prepareSauces(ch)


		Channel as Communication Path:

			In the restaurant, communication is essential. Chefs need to know when one task is completed to start the next. 
			Channels act like communication paths between chefs.
			
			// Channel for communication

						ch := make(chan string)

			For example, after chopping vegetables, a chef sends a message through the channel:
			
			// Chef 1 (goroutine) sending a message through the channel

					ch <- "Vegetables are chopped!"

			Another chef (goroutine) listens for the message and acts accordingly:
			
			// Chef 2 (goroutine) listening for a message on the channel
					message := <-ch


		Synchronization with Channels:

			Channels help synchronize tasks. For instance, a chef may not start grilling meat
			until the message "Vegetables are chopped!" is received.
			
			// Chef 2 (goroutine) waiting for a message on the channel
						message := <-ch
						if message == "Vegetables are chopped!" {
							// Chef 2 can start grilling meat
							go grillMeat(ch)
						}

		Orderly Execution:

			Channels ensure that tasks are executed in a coordinated and orderly fashion. 
			For instance, sauces may need to be ready before combining them with other dishes.
		
			// Chef 3 (goroutine) ensuring sauces are ready before proceeding

					saucesReady := <-ch
					if saucesReady {
						// Combine sauces with other dishes
						go combineDishes(ch)
					}

	In this analogy:

			Goroutines (Chefs)				: Represent concurrent tasks or functions.
			Channel (Communication Path)	: Represents the communication path between goroutines.
			Messages						: Represent information or signals sent between goroutines through the channel.


			By using channels, you ensure that the different tasks (goroutines) 
			in your program can communicate effectively, synchronize their activities, 
			and work together to accomplish the overall goal, just like chefs in a busy
			 kitchen working in harmony to deliver delicious meals.




	--------------------------------------------------------------------------------------------------------------------
*/