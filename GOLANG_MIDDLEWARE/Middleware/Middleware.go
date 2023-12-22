package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context){

	if ! (c.Request.Header.Get("Token")=="auth"){
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"token not present",
		})
		return
	}

	c.Next()

}