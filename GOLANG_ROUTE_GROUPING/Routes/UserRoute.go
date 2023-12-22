package routes

import (
	controllers "api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetRouter()*gin.Engine{

	router:= gin.Default()

	///////////////////////////////////////////////////////////////////////////
	//-----------------------THIS ROUTE IS OPEN TO ALL------------------------

	router.GET("/",controllers.OpenToAll)

	///////////////////////////////////////////////////////////////////////////
    //----------------------THIS ROUTE IS FOR ADMIN---------------------------

	Admin := router.Group("/admin")
	{
		Admin.GET("/test",controllers.WelcomeAdmin)
	}
	
	///////////////////////////////////////////////////////////////////////////
	//----------------------THIS ROUTE IS FOR USER-----------------------------	
	
	User := router.Group("/user")
	{
		User.GET("/test",controllers.WelcomeUser)
	}

	///////////////////////////////////////////////////////////////////////////

	return router

}

