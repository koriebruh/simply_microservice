package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/koriebruh/simply_microservice/dto"
	"github.com/koriebruh/simply_microservice/entity"
	"github.com/koriebruh/simply_microservice/utils"
	"gorm.io/gorm"
	"net/http"
)

type OrderController interface {
	CreateOrderController(ctx *fiber.Ctx) error
	StatusOrderController(ctx *fiber.Ctx) error
}

type OrderControllerImpl struct {
	*gorm.DB
	*validator.Validate
}

func NewOrderControllerImpl(DB *gorm.DB, validate *validator.Validate) *OrderControllerImpl {
	return &OrderControllerImpl{DB: DB, Validate: validate}
}

func (c OrderControllerImpl) CreateOrderController(ctx *fiber.Ctx) error {
	//GET REQUEST BODY
	var body dto.OrderRequest
	if err := ctx.BodyParser(&body); err != nil {
		fmt.Println("eHEHE", err)
		return utils.WebResponse(ctx, http.StatusInternalServerError, err, "internal error parsing body", nil)
	}

	//VALIDATE DULU
	if err := c.Validate.Struct(body); err != nil {
		return utils.WebResponse(ctx, http.StatusBadRequest, err, "cant accept request body bc payload not complete", nil)
	}

	fmt.Println("AMAN MASUK DATA NYA")

	// START INSERT DB AND INSERT KE KAFKA AS PRODUCER KE TOPIC  order_request MAPPING FIRST
	// INSERT INI DULU KARENA ITEMS BUTH ID ORDER
	var newOrder = entity.Order{
		Model:          gorm.Model{},
		Items:          nil,
		Amount:         body.Amount,
		PaymentMethod:  entity.BankTransfer,
		ShippingAddr:   body.ShippingAddr,
		ShippingStatus: entity.PendingShipment,
	}
	if err := c.DB.WithContext(ctx.Context()).Create(&newOrder).Error; err != nil {
		return utils.WebResponse(ctx, http.StatusInternalServerError, err, "failed to create new order", nil)
	}

	// AMBIL ID YG DI HASILKAN DAN SEND PESANAN
	var items []entity.ProductOrder
	for _, i := range body.Items {
		items = append(items, entity.ProductOrder{
			OrderID:   newOrder.ID,
			ProductID: uint(i.ProductId),
			Quantity:  i.Quantity,
		})
	}
	if err := c.DB.WithContext(ctx.Context()).Create(&items).Error; err != nil {
		return utils.WebResponse(ctx, http.StatusInternalServerError, err, "failed to save order request", nil)
	}

	// DI SINI NANTI BIKIN func  KAFKA SEND TOPIC order_request

	// RETURN SUCCESS RESPONSE
	return utils.WebResponse(ctx, http.StatusCreated, nil, "create order success", nil)

}

func (c OrderControllerImpl) StatusOrderController(ctx *fiber.Ctx) error {
	//GET REQUEST BODY
	params := ctx.Params("id")
	var orderStatus entity.Order
	if err := c.DB.WithContext(ctx.Context()).Preload("Items").Where(params).Order(&orderStatus).Error; err != nil {
		return err
	}

	//MAPPING, Ini asal dulu
	return ctx.Status(http.StatusOK).JSON(dto.WebResponse{
		Status:  "success",
		Message: "create order success",
		Data:    orderStatus,
	})

}
