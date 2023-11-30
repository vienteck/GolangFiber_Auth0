package router

import (
	"encoding/gob"
	authenticator "fiber_idle/platform/auth"
	"fiber_idle/platform/middleware"
	"fiber_idle/web/app/callback"
	"fiber_idle/web/app/index"
	"fiber_idle/web/app/login"
	"fiber_idle/web/app/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func New(auth *authenticator.Authenticator) *fiber.App {
	engine := html.New("./web/template", ".html")

	router := fiber.New(fiber.Config{
		Views: engine,
	})

	gob.Register(map[string]interface{}{})

	router.Static("/public", "web/static")

	store := session.New()

	middleware.SetupSessionStoreMiddleware(router, store)
	middleware.IsAuthenticatedMiddleware(router)

	router.Get("/", index.Handler)
	router.Get("login", login.Handler(auth))
	router.Get("callback", callback.Handler(auth))
	router.Get("user", user.Handler)
	return router
}
