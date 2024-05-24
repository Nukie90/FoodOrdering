package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/restaurant/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type RestaurantHandler struct {
	restaurantUsecase *usecase.RestaurantUsecase
}

func NewRestaurantHandler(usecase *usecase.RestaurantUsecase) *RestaurantHandler {
	return &RestaurantHandler{restaurantUsecase: usecase}
}

func (h *RestaurantHandler) CreateRestaurant(c *fiber.Ctx) error {
	var reqForm model.CreateRestaurant
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	err := h.restaurantUsecase.CreateRestaurant(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *RestaurantHandler) AdjustTable(c *fiber.Ctx) error {
	params := c.Params("name")
	var reqForm model.AdjustTable
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	reqForm.Name = params
	
	err := h.restaurantUsecase.AdjustTable(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}