package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/middleware"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
)

func Todos(controller *controllers.Todos, authService *services.Auth, app fiber.Router) fiber.Router {
	return app.Route("/todos", func(todos fiber.Router) {
		todos.Get("/", controller.Find)
		todos.Get("/:id", controller.FindOneById)
		todos.Post("/", middleware.Authenticate(authService), controller.Create)
		todos.Put("/:id", middleware.Authenticate(authService), controller.Update)
		todos.Delete("/:id", middleware.Authenticate(authService), controller.Delete)
	})
}
