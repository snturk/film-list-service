package configs

import (
	"filmListService/routes"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	routes.UserRoute(app)
}