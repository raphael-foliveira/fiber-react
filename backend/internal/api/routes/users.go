package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
)

func Users(controller *controllers.Users, app fiber.Router, authMw func(*fiber.Ctx) error) fiber.Router {
	return app.Route("/users", func(users fiber.Router) {
		users.Put("/:id", authMw, controller.Update)
		users.Get("/:id/todos", authMw, controller.FindUserTodos)
		users.Delete("/:id", authMw, controller.Delete)
	})
}
