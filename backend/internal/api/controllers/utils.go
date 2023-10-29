package controllers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/fiber-react/backend/internal/errs"
)

func parseAuthHeader(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()
	authorization, ok := headers["Authorization"]
	if !ok {
		return "", errs.HTTPError{Code: 401, Message: "Unauthorized"}
	}
	token := strings.Split(authorization[0], " ")
	if len(token) < 2 {
		return "", errs.HTTPError{Code: 401, Message: "Unauthorized"}
	}
	return token[1], nil
}
