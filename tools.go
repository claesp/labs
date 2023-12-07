package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func toolsIndex(c *fiber.Ctx) error {
	return c.Render("tools-index", fiber.Map{
		"Title": "Tools | labs.cx.se",
	})
}

func toolsFor(c *fiber.Ctx) error {
	return c.Render("tools-for", fiber.Map{
		"Title": "For | Tools | labs.cx.se",
	})
}

type ForData struct {
	InputData  string `form:"idata"`
	Template   string `form:"tmplt"`
	OutputData string `form:"odata"`
}

func toolsForExec(c *fiber.Ctx) error {
	f := new(ForData)
	if err := c.BodyParser(f); err != nil {
		return err
	}

	f.OutputData = ""

	if f.InputData != "" && f.Template != "" {
		for _, ir := range strings.Split(f.InputData, "\r\n") {
			ic := strings.Split(ir, ",")
			d := make([]interface{}, 0)
			for _, c := range ic {
				d = append(d, c)
			}
			f.OutputData = f.OutputData + fmt.Sprintf(f.Template, d...) + "\n"
		}
	}

	return c.Render("tools-for", fiber.Map{
		"Title":      "For | Tools | labs.cx.se",
		"InputData":  f.InputData,
		"Template":   f.Template,
		"OutputData": f.OutputData,
	})
}
