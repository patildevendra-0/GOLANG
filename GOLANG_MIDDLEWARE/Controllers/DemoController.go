package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c * gin.Context){

	c.JSONP(http.StatusOK,gin.H{
		"message":"welcome to api",
	})

}

func Demo(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"welcome to demo api",
	})
}