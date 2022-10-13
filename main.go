package main

import (
	"flag"
	"fmt"
	"images/database"
	"images/routes"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var (
	host    = flag.String("host", "127.0.0.1:3000", "Specifies the interface to listen")
	prefork = flag.Bool("prefork", false, "Specifies if prefork will be enabled")
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Couldn't load .env file")
	}

	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork:      *prefork,
		UnescapePath: true,
	})

	db := database.New()
	files := database.GetFiles(db)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.Register(files, app.Group("/"))

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	log.Fatal(app.Listen(*host))
}
