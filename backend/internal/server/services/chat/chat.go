package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/server"
)

func init() {
	server.RegisterHandler(fiber.MethodGet, "/hello", auth.NONE, helloWorld)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello Worl")
}