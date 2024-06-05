package handler

import (
	"foodOrder/domain/model"
	"foodOrder/internal/api/ordering/usecase"
	_ "net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderingHandler struct {
	orderingUsecase *usecase.OrderingUsecase
}

func NewOrderingHandler(usecase *usecase.OrderingUsecase) *OrderingHandler {
	return &OrderingHandler{orderingUsecase: usecase}
}

func (h *OrderingHandler) AddToCart(c *fiber.Ctx) error {
	var reqForm model.AddToCart
	if err := c.BodyParser(&reqForm); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	reqForm.UserOrder = c.Locals("guestId").(string)
	oldTableno := c.Locals("tableNo")
	reqForm.TableNo = uint8(oldTableno.(int))

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
	tableNo := c.Locals("tableNo").(int)
	cart, err := h.orderingUsecase.CartDetail(uint8(tableNo))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    cart,
	})
}

func (h *OrderingHandler) SubmitCart(c *fiber.Ctx) error {
	tableNo := c.Locals("tableNo").(int)
	err := h.orderingUsecase.SubmitCart(uint8(tableNo))
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
	allOrder, err := h.orderingUsecase.ReceiveOrder()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    allOrder,
	})

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