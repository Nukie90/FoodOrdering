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

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userHandler := uh.NewUserHandler(uu.NewUserUsecase(ur.NewUserRepo(db)))
	foodHandler := fh.NewFoodHandler(fu.NewFoodUsecase(fr.NewFoodRepo(db)))
	restaurantHandler := rh.NewRestaurantHandler(ru.NewRestaurantUsecase(rr.NewRestRepo(db)))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.Login)
	app.Get("/all", userHandler.GetAllUsers)
	app.Delete("/delete", userHandler.DeleteAll)

	staff := app.Group("/staff")
	{
		staff.Use(validating.JWTAuth(), validating.IsStaff())
		staff.Post("/add", foodHandler.CreateFood)
		staff.Get("/all", foodHandler.GetAllFoods)
		staff.Post("/restaurant", restaurantHandler.CreateRestaurant)
		staff.Put("update/:name", restaurantHandler.AdjustTable)
	}
}