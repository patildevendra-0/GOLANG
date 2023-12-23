package controllers

import "github.com/gofiber/fiber/v2"

func Welcome(c* fiber.Ctx)error{

	data := c.JSON(fiber.Map{
		"message":"welcome to demo api using fiber",
	})

	return data
}