package callback

import (
	authenticator "fiber_idle/platform/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Handler(auth *authenticator.Authenticator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session := c.Locals("session").(*session.Session)

		if c.Query("state") != session.Get("state") {
			return c.SendStatus(http.StatusBadRequest)
		}
		// Exchange an authorization code for a token.j
		token, err := auth.Exchange(c.Context(), c.Query("code"))
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		idToken, err := auth.VerifyIDToken(c.Context(), token)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		session.Set("profile", profile)

		if err := session.Save(); err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.Redirect("/user", http.StatusTemporaryRedirect)
	}
}
