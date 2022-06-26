package main

import (
	"filmListService/configs"
	"filmListService/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	repository.ConnectDB()
	configs.InitRoutes(app)

	app.Listen(":6000")
}
