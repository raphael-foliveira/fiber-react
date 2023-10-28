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
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/routes"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
	"github.com/raphael-foliveira/fiber-react/backend/internal/persistence/repositories"
)

func Start(db *sql.DB) error {
	port := os.Getenv("HTTP_PORT")
	app := fiber.New(appConfig)
	attachMiddleware(app)

	todosRepository := repositories.NewTodos(db)
	usersRepository := repositories.NewUsers(db)
	refreshTokensRepository := repositories.NewRefreshTokens(db)

	todosService := services.NewTodos(todosRepository)
	usersService := services.NewUsers(usersRepository)
	jwtService := services.NewJwt()
	authService := services.NewAuth(refreshTokensRepository, usersService, jwtService)

	todosController := controllers.NewTodos(todosService)
	usersController := controllers.NewUsers(usersService)
	authController := controllers.NewAuth(authService)

	routes.Todos(todosController, app)
	routes.Users(usersController, app)
	routes.Auth(authController, app)

	healthCheckController := controllers.NewHealthCheck()

	app.Get("/health-check", healthCheckController.HealthCheck)

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
	var fiberErr *fiber.Error
	var httpErr *errs.HTTPError
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
	}
	if errors.As(err, &httpErr) {
		code = httpErr.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).JSON(fiber.Map{
		"error":  err.Error(),
		"status": code,
	})
}
