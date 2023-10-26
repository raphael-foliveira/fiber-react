package controllers

import "github.com/gofiber/fiber/v2"

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (hc *HealthCheck) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}
