package api

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/contrib/swagger"
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
	usersService := services.NewUsers(usersRepository, todosService)
	jwtService := services.NewJwt()
	authService := services.NewAuth(refreshTokensRepository, usersService, jwtService)

	todosController := controllers.NewTodos(todosService)
	usersController := controllers.NewUsers(usersService, authService)
	authController := controllers.NewAuth(authService)
	healthCheckController := controllers.NewHealthCheck()

	app.Route("/api", func(router fiber.Router) {
		routes.Todos(todosController, authController, router)
		routes.Users(usersController, authController, router)
		routes.Auth(authController, router)
	})
	app.Get("/health-check", healthCheckController.HealthCheck)

	app.Static("/", "web")
	app.Get("/*", func(c *fiber.Ctx) error {
		fmt.Println("running")
		return c.SendFile("./web/index.html")
	})

	log.Info("Starting server on port: ", port)
	return app.Listen(fmt.Sprintf(":%v", port))
}

func attachMiddleware(app *fiber.App) {
	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.yaml",
		Path:     "/api",
		Title:    "Todo API",
	}))
	app.Use(cors.New())
	app.Use(logger.New())
}

var appConfig = fiber.Config{
	ErrorHandler: errorHandler,
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var fiberErr *fiber.Error
	var httpErr errs.HTTPError
	var conflictErr errs.ConflictError
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
	}
	if errors.As(err, &httpErr) {
		code = httpErr.Code
	}
	if errors.As(err, &conflictErr) {
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.Status(code).JSON(fiber.Map{
			"error":  err.Error(),
			"status": 409,
			"field":  conflictErr.Field,
		})
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).JSON(fiber.Map{
		"error":  err.Error(),
		"status": code,
	})
}
