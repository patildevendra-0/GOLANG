package routes

import (
	controllers "api/Controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App){

	app.Get("/",controllers.Welcome)
	// app.Post("/user/register",controllers.RegisterUser)
	// app.Get("/users",controllers.GetAllUsers)
	// app.Put("/user/:id",controllers.UpdateUser)
	// app.Delete("user/delete/:id",controllers.Deleteuser)
}