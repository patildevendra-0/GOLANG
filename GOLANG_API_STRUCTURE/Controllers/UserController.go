package controllers

import (
	models "api/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestApi(c *gin.Context){
	user:= models.Student{
		Name: "DEVENDRA",
	}

	c.JSON(http.StatusOK,user)

}

func DemoFunction(c *gin.Context){
	name:=c.Param("name")
	c.JSON(http.StatusOK,gin.H{
		"message":"Welcome "+name+" to our api which is under construction",
	})
}