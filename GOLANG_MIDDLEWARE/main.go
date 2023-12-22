package main

import routes "api/Routes"


func main(){

	router := routes.SetRouter()

	router.Run(":8080")


}