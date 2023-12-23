package main

import "github.com/gofiber/fiber/v2"


func main(){

	app := fiber.New()

	app.Get("/",Welcome)

	app.Listen(":8080")

}

func Welcome(c *fiber.Ctx)error{
	
	data:= c.JSON(fiber.Map{
		"message":"welcome to fiber test api",
	})

	return data
}