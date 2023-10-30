package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/controllers"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/middleware"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
)

func Auth(controller *controllers.Auth, authService *services.Auth, app fiber.Router) fiber.Router {
	return app.Route("/auth", func(auth fiber.Router) {
		auth.Post("/login", controller.Login)
		auth.Post("/signup", controller.Signup)
		auth.Post("/refresh-token", controller.RefreshToken)
		auth.Post("/logout", middleware.Authenticate(authService), controller.Logout)
	})
}
