package routes

import (
	"encoding/base64"
	"images/database"
	"io/ioutil"
	"strings"

	editor "images/tools"

	"github.com/gofiber/fiber/v2"
)

func GetImage(files *database.Files) fiber.Handler {
	return func(c *fiber.Ctx) error {
		file := files.Get(c.Params("name"))
		if file == nil {
			return c.Redirect("https://zzz.drylo.xyz/")
		}

		c.SendStatus(200)
		c.Set("Content-Type", file.Type)

		cX, cY, width, height, angle, resize, blur :=
			c.Query("x"),
			c.Query("y"),
			c.Query("width"),
			c.Query("height"),
			c.Query("angle"),
			c.Query("autoresize"),
			c.Query("blur")

		decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(file.Data))

		if file.Type == "image/png" &&
			(cX != "" || cY != "" || width != "" || height != "" || angle != "" || blur != "") {
			img := editor.ProcessImage(decoder, cX, cY, width, height, angle, resize, blur)
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
