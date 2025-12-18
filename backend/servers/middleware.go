package servers

import (
	"fmt"
	"task-management-backend/configs"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func (s *server) middlewareLogger(ctx fiber.Ctx) error {
	s.log.Info(fmt.Sprintf("IN: %s %s", ctx.Method(), ctx.Path()))

	startTime := time.Now()

	err := ctx.Next()

	duration := time.Since(startTime)
	s.log.Info(fmt.Sprintf("OUT: %s %d %s (Duration: %s)",
		ctx.Method(),
		ctx.Response().StatusCode(),
		ctx.Path(),
		duration.String(),
	))

	return err
}

func (s *server) Cors(cfg *configs.Config) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		corsConfig := cors.Config{
			AllowOrigins: cfg.App.Cors.AllowedOrigins,
			AllowHeaders: []string{"Content-Type", "Authorization"},
			AllowMethods: []string{"GET, POST, HEAD, PUT, DELETE, PATCH"},
		}
		return cors.New(corsConfig)(c)
	}
}
