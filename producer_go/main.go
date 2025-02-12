package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/koriebruh/simply_microservice/cfg"
	"github.com/koriebruh/simply_microservice/controller"
	"log"
)

func main() {
	app := fiber.New()
	config := cfg.GetConfig()
	validate := validator.New()
	pool, _ := cfg.GetPool(config)

	orderController := controller.NewOrderControllerImpl(pool, validate)

	app.Post("/api/orders", orderController.CreateOrderController)
	app.Get("/api/orders/status/:id", orderController.StatusOrderController) // ini id order ygy
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(fmt.Sprintf(":" + config.Server.Port)); err != nil {
		log.Fatal("Server Terminated")
	}
}

//NOTE REFACTOR NANTI PERBAIKI  RESPONSE NYA DAN JNAGAN LUPA YG CREATE SETING AMOUNT NYA
