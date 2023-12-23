package routes

import (
	controllers "api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetRouter()*gin.Engine{

	router:= gin.Default()

	router.GET("/",controllers.TestApi)
	router.GET("/user/:name",controllers.DemoFunction)

	return router

}