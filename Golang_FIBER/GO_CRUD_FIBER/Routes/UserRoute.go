package routes

import (
	controllers "api/Controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App){

	app.Get("/",controllers.Welcome)

}