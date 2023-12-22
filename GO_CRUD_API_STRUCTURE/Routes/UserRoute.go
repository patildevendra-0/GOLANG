package routes

import (
	controllers "api/Controllers"

	"github.com/gin-gonic/gin"
)


func RouterStarter()*gin.Engine{

	router := gin.Default()
	router.GET("/",controllers.Welcome)
	router.POST("/user/create",controllers.CreateNewUser)
	router.GET("/users",controllers.GetAllUsers)
    router.PUT("user/:id",controllers.UpdateUser)
	router.DELETE("user/delete/:id",controllers.DeleteUser)
	return router
}