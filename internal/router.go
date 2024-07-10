package internal

import (
	uh "foodOrder/internal/api/user/handler"
	ur "foodOrder/internal/api/user/repository"
	uu "foodOrder/internal/api/user/usecase"
	"foodOrder/internal/api/validating"

	fh "foodOrder/internal/api/food/handler"
	fr "foodOrder/internal/api/food/repository"
	fu "foodOrder/internal/api/food/usecase"

	rh "foodOrder/internal/api/restaurant/handler"
	rr "foodOrder/internal/api/restaurant/repository"
	ru "foodOrder/internal/api/restaurant/usecase"

	oh "foodOrder/internal/api/ordering/handler"
	or "foodOrder/internal/api/ordering/repository"
	ou "foodOrder/internal/api/ordering/usecase"

	ph "foodOrder/internal/api/payment/handler"
	pr "foodOrder/internal/api/payment/repository"
	pu "foodOrder/internal/api/payment/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userHandler := uh.NewUserHandler(uu.NewUserUsecase(ur.NewUserRepo(db)))
	foodHandler := fh.NewFoodHandler(fu.NewFoodUsecase(fr.NewFoodRepo(db)))
	restaurantHandler := rh.NewRestaurantHandler(ru.NewRestaurantUsecase(rr.NewRestRepo(db)))
	orderingHandler := oh.NewOrderingHandler(ou.NewOrderingUsecase(or.NewOrderingRepo(db)))
	paymentHandler := ph.NewPaymentHandler(pu.NewPaymentUsecase(pr.NewPaymentRepo(db)))


	//new route
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Post("/register", userHandler.RegisterUser)
			v1.Post("/login", userHandler.Login)
			v1.Get("/users", userHandler.GetAllUsers)
			v1.Get("/foods", foodHandler.GetAllFoods)

			staff := v1.Group("/staff")
			{
				staff.Use(validating.JWTAuth(), validating.IsStaff())
				staff.Post("/foods", foodHandler.CreateFood)
				staff.Post("/table-inits", restaurantHandler.InitialTable)
				staff.Get("/tables", restaurantHandler.GetAllTable)
				staff.Post("/tables", restaurantHandler.GiveCustomerTable)
				staff.Post("/payments", paymentHandler.CreatePayment)
				staff.Get("/orders", orderingHandler.ReceiveOrder)
				staff.Put("/orders", orderingHandler.UpdateOrder)
			}

			cooker := v1.Group("/cooker")
			{
				cooker.Use(validating.JWTAuth(), validating.IsCooker())
				cooker.Get("/orders", orderingHandler.ReceiveOrder)
				cooker.Post("/robots", orderingHandler.SendRobot)
			}

			customer := v1.Group("/customer")
			{
				customer.Post("/carts/:tableID", orderingHandler.AddToCart)
				customer.Get("/carts/:tableID", orderingHandler.GetCart)
				customer.Post("/orders/:tableID", orderingHandler.SubmitCart)
				customer.Get("/robots/:tableNo", orderingHandler.ReceiveRobot)
				customer.Post("/payments/:id", paymentHandler.PayBill)
			}
		}
	}
}

//new code
