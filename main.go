package main

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"main/logger"
	"main/shared"
	"time"
)

func main() {
	logger.Info("Starting application", zap.String("version", "1.0.0"))
	shared.SetupEnv(".")
	logger.Info("Application initialized successfully")

	app := fiber.New()

	app.Get(
		"/", func(c fiber.Ctx) error {
			return c.JSON(
				fiber.Map{
					"status": "ok",
					"slogan": "Wow, so this is what it's like to be on the internet!",
					"date":   time.Now().Format("2006-01-02 15:04:05"),
				},
			)
		},
	)

	app.Get(
		"/health", func(c fiber.Ctx) error {
			logger.Info("/health called")
			return c.JSON(
				fiber.Map{
					"status": "ok",
					"slogan": "All aboard the railway express!",
					"date":   time.Now().Format("2006-01-02 15:04:05"),
				},
			)
		},
	)

	port := shared.Port()
	logger.Info("Starting server", zap.String("port", port))
	logger.Fatal("Server failed to start", zap.Error(app.Listen(":"+port)))
}
