package handler

import (
	"github.com/betterde/template/fiber/internal/response"
	"github.com/gofiber/fiber/v3"
)

func HealthCheck(ctx fiber.Ctx) error {
	return response.Success(ctx)
}
