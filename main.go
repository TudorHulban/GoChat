package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// Create a new engine
	engine := html.New("./views", ".gohtml")

	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", handler)
	app.Get("/p", handlerUser)

	log.Fatal(app.Listen(":3000"))
}

func handler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Person": "John",
	})
}

func handlerUser(c *fiber.Ctx) error {
	u := User{
		Name: "John",
	}

	return c.Render("index", u)
}
