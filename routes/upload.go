package routes

import (
	"encoding/base64"
	"images/database"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(files *database.Files) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Redirect("https://zzz.drylo.xyz/")
		}

		contents, err := file.Open()
		if err != nil {
			return c.Redirect("https://zzz.drylo.xyz/")
		}

		defer contents.Close()

		bytes, err := ioutil.ReadAll(contents)
		if err != nil {
			return c.Redirect("https://zzz.drylo.xyz/")
		}

		f := files.Upload(base64.StdEncoding.EncodeToString(bytes), file.Header.Get("Content-Type"))
		return c.JSON(fiber.Map{
			"id":   f.ID,
			"name": f.Name,
		})
	}
}
