package controllers

import "github.com/gofiber/fiber/v2"

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

// HealthCheck godoc
// @Summary HealthCheck
// @Description HealthCheck
// @Tags health_check
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /health_check [get]
func (hc *HealthCheck) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}
