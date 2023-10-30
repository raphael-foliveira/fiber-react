package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
)

func Todos(controller *controllers.Todos, app fiber.Router, authMw func(*fiber.Ctx) error) fiber.Router {
	return app.Route("/todos", func(todos fiber.Router) {
		todos.Get("/", controller.Find)
		todos.Get("/:id", controller.FindOneById)
		todos.Post("/", authMw, controller.Create)
		todos.Put("/:id", authMw, controller.Update)
		todos.Delete("/:id", authMw, controller.Delete)
	})
}
