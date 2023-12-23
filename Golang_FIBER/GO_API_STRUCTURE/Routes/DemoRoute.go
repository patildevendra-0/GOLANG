package routes

import (
	controllers "api/Controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(c *fiber.App){

	c.Get("/",controllers.Welcome)

}