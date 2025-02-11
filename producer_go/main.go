package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Server Terminated")
	}
}

func CrateOrderHandler(ctx *fiber.Ctx) error {

}

func OrderStatusHandler(ctx *fiber.Ctx) error {

}
