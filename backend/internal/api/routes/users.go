package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/middleware"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
)

func Users(controller *controllers.Users, authService *services.Auth, app fiber.Router) fiber.Router {
	return app.Route("/users", func(users fiber.Router) {
		users.Get("/", controller.Find)
		users.Get("/:id", controller.FindOneById)
		users.Post("/", controller.Create)
		users.Put("/:id", controller.Update)
		users.Get("/:id/todos", middleware.Authenticate(authService), controller.FindUserTodos)
	})
}
