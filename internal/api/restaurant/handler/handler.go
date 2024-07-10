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

func (h *RestaurantHandler) InitialTable(c *fiber.Ctx) error {
	var reqForm model.InitialTable
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	err := h.restaurantUsecase.InitialTable(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *RestaurantHandler) GetAllTable(c *fiber.Ctx) error {
	tables, err := h.restaurantUsecase.GetAllTable()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(tables)
}

func (h *RestaurantHandler) GiveCustomerTable(c *fiber.Ctx) error {
	var reqForm model.GiveTable
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	ID, err := h.restaurantUsecase.GiveCustomerTable(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"PereferenceID": ID.String(),
	})


}

func (h *RestaurantHandler) CheckHistory(c *fiber.Ctx) error {
	var reqForm model.CheckHistory

	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	history, err := h.restaurantUsecase.CheckHistory(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(history)
}