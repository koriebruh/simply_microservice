package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/koriebruh/simply_microservice/cfg"
	"log"
)

func main() {
	app := fiber.New()
	config := cfg.GetConfig()
	cfg.GetPool(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(fmt.Sprintf(":" + config.Server.Port)); err != nil {
		log.Fatal("Server Terminated")
	}
}
