package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/khu-dev/khumu-club/config"
	"github.com/khu-dev/khumu-club/http"
	"github.com/khu-dev/khumu-club/repository"
	"github.com/khu-dev/khumu-club/service"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Connected with database
	client := repository.Connect()
	clubService := &service.ClubService{DB: client}
	clubHandler := &http.ClubHandler{ClubService: clubService}

	// Create fiber app
	app := fiber.New(fiber.Config{})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	// JWT Middleware
	app.Use(http.NewJWTMiddleware())

	// Create a /api/v1 endpoint
	apiv1 := app.Group("/api")

	// Bind handlers
	apiv1.Post("/clubs", clubHandler.CreateClub)
	apiv1.Get("/clubs", clubHandler.ListClub)
	//v1.Post("/users", handlers.UserCreate)

	// Setup static files
	//app.Static("/", "./static/public")

	// Handle not founds
	app.Use(http.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", config.Config.Port))) // go run app.go -port=:3000

}
