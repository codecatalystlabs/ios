package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// AuthRequired is a middleware that checks if the user is authenticated
func AuthRequired(c *fiber.Ctx) error {
	store := c.Locals("store").(*session.Store)
	sess, err := store.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}

	// Check if user is authenticated
	if sess.Get("authenticated") != true {
		return c.Redirect("/login")
	}

	return c.Next()
}
