package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Handler(c *fiber.Ctx) error {
	session := c.Locals("session").(*session.Session)
	profile := session.Get("profile")

	return c.Render("user", profile)
}
