package api

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
)

func Start(db *sql.DB) error {
	port := os.Getenv("HTTP_PORT")
	app := fiber.New(appConfig)
	healthCheckController := controllers.NewHealthCheck()
	app.Get("/health-check", healthCheckController.HealthCheck)

	app.Group("/api")

	log.Info("Starting server on port: ", port)
	return app.Listen(fmt.Sprintf(":%v", port))
}

func attachMiddleware(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
}

var appConfig = fiber.Config{
	ErrorHandler: errorHandler,
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).JSON(fiber.Map{
		"error":  err.Error(),
		"status": code,
	})
}
