package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/guestUser/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type GuestHandler struct {
	guestUsecase *usecase.GuestUsecase
}

func NewGuestHandler(usecase *usecase.GuestUsecase) *GuestHandler {
	return &GuestHandler{guestUsecase: usecase}
}

func (h *GuestHandler) EnterTable(c *fiber.Ctx) error {
    tableNoStr := c.Params("id")
    tableNo, err := strconv.Atoi(tableNoStr)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid table number",
        })
    }

    table := &model.EnterTable{
        TableNo: uint8(tableNo),
    }
    err = h.guestUsecase.EnterTable(table)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    c.Locals("tableNo", tableNoStr)
    return c.Next()
}
