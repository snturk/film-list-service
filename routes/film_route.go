package routes

import (
	"filmListService/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/film/getAll", controllers.GetAllFilms)
	app.Post("/film/add", controllers.AddFilm)
	app.Post("/film/delete", controllers.DeleteFilm)
}