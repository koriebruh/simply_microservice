package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OrderController interface {
	CreateOrderController(ctx *fiber.Ctx) error
	StatusOrderController(ctx *fiber.Ctx) error
}

type OrderControllerImpl struct {
	*gorm.DB
}

func NewOrderControllerImpl(DB *gorm.DB) *OrderControllerImpl {
	return &OrderControllerImpl{DB: DB}
}

func (c OrderControllerImpl) CreateOrderController(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (c OrderControllerImpl) StatusOrderController(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
