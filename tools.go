package main

import "github.com/gofiber/fiber/v2"

func toolsIndex(c *fiber.Ctx) error {
	return c.Render("tools-index", fiber.Map{
		"Title": "Tools | labs.cx.se",
	})
}
