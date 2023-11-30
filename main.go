package main

import (
	authenticator "fiber_idle/platform/auth"
	"fiber_idle/platform/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}
	app := router.New(auth)
	log.Fatal(app.Listen(":3000"))
}
