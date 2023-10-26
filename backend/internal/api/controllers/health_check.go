package controllers

import "github.com/gofiber/fiber/v2"

type healthCheck struct{}

func NewHealthCheck() *healthCheck {
	return &healthCheck{}
}

func (hc *healthCheck) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}
