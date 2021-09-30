package routes

import (
	"fmt"
	editor "images/tools"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	},
}

func ProcessImage(c *fiber.Ctx) error {
	url, err := url.Parse(c.Query("url"))
	if err != nil {
		return c.Redirect("https://zzz.drylo.xyz/")
	}

	cX, cY, width, height, angle, resize, blur :=
		c.Query("x"),
		c.Query("y"),
		c.Query("width"),
		c.Query("height"),
		c.Query("angle"),
		c.Query("autoresize"),
		c.Query("blur")

	if cX == "" && cY == "" && width == "" && height == "" && angle == "" && blur == "" {
		return c.Redirect("https://zzz.drylo.xyz/")
	}

	res, err := client.Get(url.String())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return c.Redirect("https://zzz.drylo.xyz/")
	}

	if res.Header.Get("Content-Type") != "image/png" {
		return c.Redirect("https://zzz.drylo.xyz/")
	}

	img := editor.ProcessImage(res.Body, cX, cY, width, height, angle, resize, blur)
	if img == nil {
		return c.Redirect("https://zzz.drylo.xyz/")
	}

	return c.Send(img)
}
