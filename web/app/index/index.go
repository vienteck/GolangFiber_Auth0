package index

import "github.com/gofiber/fiber/v2"

func Handler(c *fiber.Ctx) error {

	return c.Render("index", "")
}
