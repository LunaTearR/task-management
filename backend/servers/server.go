package servers

import (
	"fmt"
	"strings"
	"task-management-backend/configs"
	"task-management-backend/pkg"
	"runtime/debug"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type server struct {
	App *fiber.App
	cfg *configs.Config
	log pkg.AppLog
}

func NewServer(config *configs.Config, log pkg.AppLog) *server {
	fiberCfg := fiber.Config{
		AppName:         config.App.Name,
		ProxyHeader:     "X-Forwarded-For",
		ReadBufferSize:  1024 * 1024 * 15,
		WriteBufferSize: 1024 * 1024 * 15,
	}

	return &server{
		App: fiber.New(fiberCfg),
		cfg: config,
	}
}

func (s *server) Start() {
	s.setupMiddleware()

	s.App.Get("*", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Route not found",
			"path":    c.Path(),
		})
	})

	fiberConnURL := fmt.Sprintf(":%s", s.cfg.App.Port)
	if strings.ToLower(s.cfg.App.Env) == "development" {
		fiberConnURL = "0.0.0.0:" + s.cfg.App.Port
	}
	fmt.Println("fiberConnURL >>>", fiberConnURL)
	if err := s.App.Listen(fiberConnURL); err != nil {
		panic(err.Error())
	}
}

func (s *server) setupMiddleware() {
	s.App.Use(s.Cors(s.cfg))

	s.App.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(ctx fiber.Ctx, e interface{}) {
			s.log.Error("", fmt.Errorf("PANIC: %s %d %s {error %v\nStack trace:\n%s}",
				ctx.Method(),
				ctx.Response().StatusCode(),
				ctx.Path(),
				e,
				debug.Stack(),
			))
		},
	}))
	
	s.App.Use(s.middlewareLogger)
	s.App.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}
