package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	if c.Get("Authorization") != os.Getenv("SECRET_KEY") {
		return c.Redirect("/")
	}

	return c.Next()
}
