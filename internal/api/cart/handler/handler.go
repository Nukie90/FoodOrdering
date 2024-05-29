package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/cart/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	cartUsecase *usecase.CartUsecase
}

func NewCartHandler(usecase *usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: usecase}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	var reqForm model.AddToCart
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	reqForm.UserOrder = c.Locals("guestId").(string)
	oldTableno := c.Locals("tableNo")
	reqForm.TableNo = uint8(oldTableno.(int))

	err := h.cartUsecase.AddToCart(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}