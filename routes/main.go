package routes

import (
	"images/database"
	"images/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Register(files *database.Files, router fiber.Router) {
	router.Post("/lets_upload_a_new_file", middlewares.CheckAuth, UploadFile(files))
	router.Get("/:name", GetImage(files))
}
