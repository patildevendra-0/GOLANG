package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()

	router.GET("",test_api)
	router.GET("/welcome",welcome_test)
	router.Run(":8080")
}

func test_api(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"SHREE GANESH",
	})
}

func welcome_test(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"WELCOME TO API.....UNDER CONSTRUCTION..........",
	})
}