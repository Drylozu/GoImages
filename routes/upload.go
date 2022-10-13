package routes

import (
	"encoding/base64"
	"fmt"
	"images/database"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(files *database.Files) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Printf("err: %#v\n", err)
			return c.Redirect("/")
		}

		contents, err := file.Open()
		if err != nil {
			fmt.Printf("err: %#v\n", err)
			return c.Redirect("/")
		}

		defer contents.Close()

		bytes, err := ioutil.ReadAll(contents)
		if err != nil {
			fmt.Printf("err: %#v\n", err)
			return c.Redirect("/")
		}

		f := files.Upload(base64.StdEncoding.EncodeToString(bytes), file.Header.Get("Content-Type"))
		return c.JSON(fiber.Map{
			"id": f.ID,
		})
	}
}
