package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/authUser/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: usecase}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var reqForm model.RegisterUser
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	err := h.userUsecase.RegisterUser(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    users,
	})
}

func (h *UserHandler) DeleteAll(c *fiber.Ctx) error {
	err := h.userUsecase.DeleteAll()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var reqForm model.LoginUser
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	token, err := h.userUsecase.Login(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//set bearer to postman
	c.Set("Authorization", "Bearer "+token)

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"token":    token,
	})
}