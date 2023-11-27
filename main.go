package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var (
	APP_NAME   = "labs"
	APP_CONFIG Config
	MAJOR      = 0
	MINOR      = 0
	REVISION   = 231126
)

func version() string {
	return fmt.Sprintf("%s v.%d.%d.%d", APP_NAME, MAJOR, MINOR, REVISION)
}

func main() {
	log.Println(version())

	APP_CONFIG = loadConfig()

	engine := html.New("./views", ".html")
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		AppName:               APP_NAME,
		DisableStartupMessage: true,
		Views:                 engine,
		ViewsLayout:           "layouts/main",
	})

	app.Get("/", rootIndex)

	app.Listen(fmt.Sprintf(":%d", APP_CONFIG.Port))
}
