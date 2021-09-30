package routes

import (
	"images/database"

	"github.com/gofiber/fiber/v2"
)

func Register(files *database.Files, router fiber.Router) {
	router.Get("/image", ProcessImage)
	router.Post("/lets_upload_a_new_file", UploadFile(files))
	router.Get("/:name", GetImage(files))
}
