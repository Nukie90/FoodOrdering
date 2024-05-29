package internal

import (
	uh "foodOrder/internal/api/authUser/handler"
	ur "foodOrder/internal/api/authUser/repository"
	uu "foodOrder/internal/api/authUser/usecase"
	"foodOrder/internal/api/validating"
	"strconv"

	fh "foodOrder/internal/api/food/handler"
	fr "foodOrder/internal/api/food/repository"
	fu "foodOrder/internal/api/food/usecase"

	rh "foodOrder/internal/api/restaurant/handler"
	rr "foodOrder/internal/api/restaurant/repository"
	ru "foodOrder/internal/api/restaurant/usecase"

	gh "foodOrder/internal/api/guestUser/handler"
	gr "foodOrder/internal/api/guestUser/repository"
	gu "foodOrder/internal/api/guestUser/usecase"

	ch "foodOrder/internal/api/cart/handler"
	cr "foodOrder/internal/api/cart/repository"
	cu "foodOrder/internal/api/cart/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userHandler := uh.NewUserHandler(uu.NewUserUsecase(ur.NewUserRepo(db)))
	foodHandler := fh.NewFoodHandler(fu.NewFoodUsecase(fr.NewFoodRepo(db)))
	restaurantHandler := rh.NewRestaurantHandler(ru.NewRestaurantUsecase(rr.NewRestRepo(db)))
	guestHandler := gh.NewGuestHandler(gu.NewGuestUsecase(gr.NewGuestRepo(db)))
	cartHandler := ch.NewCartHandler(cu.NewCartUsecase(cr.NewCartRepo(db)))

	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.Login)
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/menu", foodHandler.GetAllFoods)
	
	staff := app.Group("/staff")
	{
		staff.Use(validating.JWTAuth(), validating.IsStaff())
		staff.Post("/add", foodHandler.CreateFood)
		staff.Post("/restaurant", restaurantHandler.CreateRestaurant)
		staff.Put("update/:name", restaurantHandler.AdjustTable)
	}
	
	guest := app.Group("/:id")
	{
		guest.Use(guestHandler.EnterTable)
		guest.Get("/table", func(c *fiber.Ctx) error {
			tableNo := c.Locals("tableNo")
			return c.JSON(fiber.Map{
				"message": "Welcome to table " + strconv.Itoa(tableNo.(int)),
				"guestId": c.Locals("guestId"),
			})
		})
		guest.Post("/addtocart", cartHandler.AddToCart)
	}
	
	
}