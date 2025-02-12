package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/koriebruh/simply_microservice/cfg"
	"github.com/koriebruh/simply_microservice/delivery"
	"github.com/koriebruh/simply_microservice/dto"
	"github.com/koriebruh/simply_microservice/entity"
	"github.com/koriebruh/simply_microservice/utils"
	"gorm.io/gorm"
	"log"
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
	cfg.GetConfig()
	newOrder.Items = items
	if err := delivery.OrderKafkaProducer(cfg.GetConfig(), "order_created", newOrder); err != nil {
		log.Println("failed publish message")
	}

	// RETURN SUCCESS RESPONSE
	return utils.WebResponse(ctx, http.StatusCreated, nil, "create order success", nil)

}

func (c OrderControllerImpl) StatusOrderController(ctx *fiber.Ctx) error {
	//GET REQUEST BODY
	params := ctx.Params("id")
	var orderStatus entity.Order
	if err := c.DB.WithContext(ctx.Context()).Where("id = ?", params).Preload("Items").First(&orderStatus).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.WebResponse(ctx, http.StatusNotFound, err, "error data not found", nil)
		}
		return utils.WebResponse(ctx, http.StatusInternalServerError, err, "got error", nil)
	}

	fmt.Printf("paramSend %v | data:%v", params, orderStatus)

	//MAPPING DATA
	var items []dto.Product
	for _, i := range orderStatus.Items {
		items = append(items, dto.Product{
			ProductId: int64(i.ProductID),
			Quantity:  i.Quantity,
		})
	}

	result := dto.OrderStatusResponse{
		Items:          items,
		Amount:         orderStatus.Amount,
		PaymentMethod:  string(orderStatus.PaymentMethod),
		ShippingAddr:   orderStatus.ShippingAddr,
		ShippingStatus: string(orderStatus.ShippingStatus),
	}

	//MAPPING, Ini asal dulu
	return utils.WebResponse(ctx, http.StatusOK, nil, "get order status success", result)

}
