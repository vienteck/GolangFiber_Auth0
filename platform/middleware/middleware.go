package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const key = "session"
const profile = "profile"

func IsAuthenticatedMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {

		if skipAuthentication(c.Path()) {
			return c.Next()
		}

		session := c.Locals(key).(*session.Session)
		temp := session.Get("profile")
		if temp == nil {
			fmt.Printf("No User profile found. Redirecting to login page. Source : %v\n", c.Path())
			return c.Redirect("/")
		} else {
			c.Locals("profile", profile)
		}
		user := session.Get("user")
		c.Locals("user", user)

		return c.Next()
	})
}

func skipAuthentication(userpath string) bool {
	//if you don't want to check for authentication add the endpoint to the paths slice
	paths := []string{"/login", "/", "/callback", "/logout", "/about", "/index"}
	for _, path := range paths {
		if userpath == path {
			return true
		}
	}
	return false
}
func SetupSessionStoreMiddleware(app *fiber.App, store *session.Store) {
	app.Use(func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		c.Locals(key, sess)
		return c.Next()
	})
}
