package handler

import (
	"fmt"
	"foodOrder/domain/model"
	"foodOrder/internal/api/ordering/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

type OrderingHandler struct {
	orderingUsecase *usecase.OrderingUsecase
}

func NewOrderingHandler(usecase *usecase.OrderingUsecase) *OrderingHandler {
	return &OrderingHandler{orderingUsecase: usecase}
}

func (h *OrderingHandler) AddToCart(c *fiber.Ctx) error {
	params := c.Params("tableID")

	var reqForm model.AddToCart

	reqForm.TableID = params

	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	err := h.orderingUsecase.AddToCart(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message1": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *OrderingHandler) GetCart(c *fiber.Ctx) error {
	tableID := c.Params("tableID")
	
	tableIDString := ulid.MustParse(tableID).String()

	cart, err := h.orderingUsecase.GetCart(tableIDString)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(cart)
}

func (h *OrderingHandler) SubmitCart(c *fiber.Ctx) error {
	tableID := c.Params("tableID")
	
	tableIDString := ulid.MustParse(tableID).String()

	fmt.Println(tableIDString)

	err := h.orderingUsecase.SubmitCart(tableIDString)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *OrderingHandler) ReceiveOrder(c *fiber.Ctx) error {
	allOrders, err := h.orderingUsecase.ReceiveOrder()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(allOrders)
}

func (h *OrderingHandler) SendRobot(c *fiber.Ctx) error {
	var reqForm model.SendRobotRequest

	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	tableNo, err := h.orderingUsecase.SendRobot(&reqForm)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Robot is on the way!",
		"tableNo": tableNo,
	})
}

func (h *OrderingHandler) ReceiveRobot(c *fiber.Ctx) error {
	tableNo := c.Locals("tableNo").(int)
	err := h.orderingUsecase.ReceiveRobot(uint8(tableNo))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Robot has arrived!",
	})
}