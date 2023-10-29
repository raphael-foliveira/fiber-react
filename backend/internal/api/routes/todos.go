package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
)

func Todos(controller *controllers.Todos, authController *controllers.Auth, app fiber.Router) fiber.Router {
	return app.Route("/todos", func(todos fiber.Router) {
		todos.Get("/", controller.Find)
		todos.Get("/:id", controller.FindOneById)
		todos.Post("/", authController.Authenticate, controller.Create)
		todos.Put("/:id", authController.Authenticate, controller.Update)
		todos.Delete("/:id", authController.Authenticate, controller.Delete)
	})
}
