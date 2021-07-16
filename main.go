package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/khu-dev/khumu-club/adapter/slack"
	"github.com/khu-dev/khumu-club/config"
	"github.com/khu-dev/khumu-club/http"
	"github.com/khu-dev/khumu-club/repository"
	"github.com/khu-dev/khumu-club/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		DisableQuote:  true,
		ForceColors:   true,
		// line을 깔끔하게 보여줌.
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			//filename := strings.Replace(f.File, workingDir + "/", "", -1)
			filename := f.File
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		FullTimestamp:   true,
		TimestampFormat: "2006/01/03 15:04:05",
	})
}

func main() {
	// Connected with database
	client := repository.Connect()
	slackAdapter := slack.NewSlackAdapter(config.Config.Slack.Token, config.Config.Slack.Channel)
	clubService := &service.ClubService{DB: client, SlackAdapter: slackAdapter}
	clubHandler := &http.ClubHandler{ClubService: clubService}
	// Create fiber app
	app := fiber.New(fiber.Config{})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	// JWT Middleware
	app.Use(http.NewJWTMiddleware())

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.SendString("healthy")
	})
	// Create a /api/v1 endpoint
	api := app.Group("/api")

	// Bind handlers
	api.Post("/clubs", clubHandler.CreateClub)
	api.Get("/clubs", clubHandler.ListClubs)
	api.Post("/clubs/add-request", clubHandler.ClubAddRequest)
	api.Post("/clubs/modify-request", clubHandler.ClubModifyRequest)
	api.Post("/club-categories", clubHandler.ListCategories)

	//v1.Post("/users", handlers.UserCreate)

	// Setup static files
	//app.Static("/", "./static/public")

	// Handle not founds
	app.Use(http.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", config.Config.Port))) // go run app.go -port=:3000

}
