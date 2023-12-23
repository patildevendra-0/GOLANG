package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeAdmin(c *gin.Context){

	c.JSON(http.StatusOK,gin.H{
		"message":"welcome to api Admin",
	})

}

func WelcomeUser(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"welcome to api user",
	})
}


func OpenToAll(c * gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"WELCOME TO API ..............UNDER CONSTRUCTION..",
	})
}