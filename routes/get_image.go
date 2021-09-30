package routes

import (
	"encoding/base64"
	"images/database"
	"io/ioutil"
	"strings"

	editor "images/tools"

	"github.com/gofiber/fiber/v2"
)

type Queries struct {
	Width  int `query:"width,omitempty"`
	Height int `query:"height,omitempty"`
}

func GetImage(files *database.Files) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file := files.Get(c.Params("name"))
		if file == nil {
			return c.Redirect("https://zzz.drylo.xyz/")
		}

		c.SendStatus(200)
		c.Set("Content-Type", file.Type)

		var q Queries

		if err := c.QueryParser(&q); err != nil {
			q = Queries{}
		}

		decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(file.Data))

		if file.Type == "image/png" && (q.Width > 0 || q.Height > 0) {
			img := editor.ProcessImage(decoder, q.Width, q.Height)
			if img == nil {
				return c.Redirect("https://zzz.drylo.xyz/")
			}

			return c.Send(img)
		}

		bytes, err := ioutil.ReadAll(decoder)
		if err != nil {
			return c.Redirect("https://zzz.drylo.xyz/")
		}

		return c.Send(bytes)
	}
}
