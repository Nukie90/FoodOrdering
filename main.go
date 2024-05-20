package main

import (
	"flag"
	_ "fmt"
	"foodOrder/infrastructure"
	"foodOrder/internal"
	"foodOrder/internal/api/validating"
	"log"
	_ "net/http"
	"strconv"
	_ "time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	envFlag := flag.String("env", "common", "a string")

	flag.Parse()

	configDetail, err := infrastructure.LoadConfig(*envFlag)
	if err != nil {
		log.Fatal(err)
	}

	dbConfig := infrastructure.NewGormConfig(configDetail)
	db, err := dbConfig.Connection()
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.AutoMigrate(db)

	app := fiber.New()

	validating.SetupMiddleware(app)
	
	serverConfig := internal.NewServerConfig(configDetail)

	serverConfigStr := serverConfig.Host + ":" + strconv.Itoa(serverConfig.Port)
	app.Listen(serverConfigStr)
}