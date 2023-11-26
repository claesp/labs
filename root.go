package main

import "github.com/gofiber/fiber/v2"

func rootIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
