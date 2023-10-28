package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
)

func Todos(controller *controllers.Todos, app fiber.Router) fiber.Router {
	return app.Route("/todos", func(todos fiber.Router) {
		todos.Get("/", controller.Find)
		todos.Get("/:id", controller.FindOneById)
		todos.Post("/", controller.Create)
		todos.Put("/:id", controller.Update)
		todos.Delete("/:id", controller.Delete)
	})
}
