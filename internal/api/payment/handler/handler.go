package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/payment/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	paymentUsecase *usecase.PaymentUsecase
}

func NewPaymentHandler(usecase *usecase.PaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUsecase: usecase}
}

func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	var reqForm model.CreatePayment
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	bill, err := h.paymentUsecase.CreatePayment(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    bill,
	})
}

