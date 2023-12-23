package main

import routes "api/Routes"

func main(){

	router:= routes.RouterStarter()

	router.Run(":8080")

}