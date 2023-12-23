package main

import (
	routes "api/Routes"

	"github.com/gofiber/fiber/v2"
)


func main(){

	app := fiber.New()

	routes.SetRoutes(app)

	app.Listen(":8080")
}