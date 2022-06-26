package controllers

import (
	"context"
	"filmListService/models"
	"filmListService/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetAllFilms(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	films := service.GetAllFilms(ctx)

	return c.JSON(films)
}

func AddFilm(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var film models.Film

	if err := c.BodyParser(&film); err != nil {
		return c.JSON(err)
	}

	service.InsertFilm(ctx, film)

	return c.JSON(service.GetAllFilms(ctx))
}

func DeleteFilm(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var film models.Film

	if err := c.BodyParser(&film); err != nil {
		return c.JSON(err)
	}

	service.DeleteFilm(ctx, film)

	return c.JSON(service.GetAllFilms(ctx))
}