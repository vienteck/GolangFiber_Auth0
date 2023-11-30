package login

import (
	"crypto/rand"
	"encoding/base64"
	authenticator "fiber_idle/platform/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Handler(auth *authenticator.Authenticator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		state, err := generateRandomState()
		if err != nil {
			c.SendStatus(http.StatusInternalServerError)
		}

		//save the state inside the session
		session := c.Locals("session").(*session.Session)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		c.Redirect(auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
		return nil
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
