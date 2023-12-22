package routes

import (
	controllers "api/Controllers"
	// middleware "api/Middleware"
	"github.com/gin-gonic/gin"
)


func SetRouter()*gin.Engine{

	router := gin.Default()


	///////////////////////////////////////////////////////////////////////////
	//----------- IF I WANT TO USE MIDDLE WARE TO ALL ROUTES THEN I USE THIS---
	/*

		router.Use(middleware.Authenticate)     							------ directly apply to all routes

	*/
    ///////////////////////////////////////////////////////////////////////////


	////////////////////////////////////////////////////////////////////////////
	//---------- IF I WANT USE FOR SPECIFIC ROUTE GROUP THEN USE THIS----------
	/*

		Admin := router.Group("/admin",middleware.Authenticate)       	 			---- apply to group routes
		{
			Admin.GET("/welcome",controllers.Welcome)
		}

	*/
	////////////////////////////////////////////////////////////////////////////


	////////////////////////////////////////////////////////////////////////////
	//-----------IF I WANT TO USE FOR SPECIFIC ROUTE IN MULTIPLE ROUTE THEN-----
	/*
	
		router.GET("/welcome",middleware.Authenticate,controllers.Welcome)   		--- apply to specific route
		router.GET("/demo",controllers.Demo)

	*/

	////////////////////////////////////////////////////////////////////////////
	//----------ROUTES WITHOUT ANY MIDDLEWARE----------------------------------

	router.GET("/welcome",controllers.Welcome)                   // 			------ routes without middleware
	router.GET("/demo",controllers.Demo)

	return router

}