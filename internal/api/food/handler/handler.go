package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/food/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type FoodHandler struct {
	foodUsecase *usecase.FoodUsecase
}

func NewFoodHandler(usecase *usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{foodUsecase: usecase}
}

func (h *FoodHandler) CreateFood(c *fiber.Ctx) error {
	var reqForm model.CreateFood
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	err := h.foodUsecase.CreateFood(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *FoodHandler) GetAllFoods(c *fiber.Ctx) error {
	foods, err := h.foodUsecase.GetAllFoods()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(foods)
}