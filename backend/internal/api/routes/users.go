package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
)

func Users(controller *controllers.Users, authController *controllers.Auth, app fiber.Router) fiber.Router {
	return app.Route("/users", func(users fiber.Router) {
		users.Get("/", controller.Find)
		users.Get("/:id", controller.FindOneById)
		users.Post("/", controller.Create)
		users.Put("/:id", controller.Update)
		users.Get("/:id/todos", authController.Authenticate, controller.FindUserTodos)
	})
}
