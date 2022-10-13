package routes

import (
	"images/database"
	"images/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Register(files *database.Files, router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("hi"))
	})

	router.Post("/upload", middlewares.CheckAuth, UploadFile(files))
	router.Get("/:id", GetImage(files))
}
