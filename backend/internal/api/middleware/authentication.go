package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/api/services"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
)

func Authorize(authService *services.Auth) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, err := parseAuthHeader(c)
		if err != nil {
			return err
		}
		fmt.Println(token)
		user, err := authService.Authorize(token)
		if err != nil {
			return err
		}
		c.Locals("user", user)
		return c.Next()
	}
}

func parseAuthHeader(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()
	authorization, ok := headers["Authorization"]
	if !ok {
		return "", &errs.HTTPError{Code: 401, Message: "Unauthorized"}
	}
	token := strings.Split(authorization[0], " ")
	if len(token) < 2 {
		return "", &errs.HTTPError{Code: 401, Message: "Unauthorized"}
	}
	return token[1], nil
}
