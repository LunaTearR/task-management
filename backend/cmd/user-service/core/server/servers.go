package server

import (
	"fmt"
	"strings"
	"task-management-user-service/configs"
	userctl "task-management-user-service/core/controllers"
	userrepo "task-management-user-service/core/repositories"
	usersrv "task-management-user-service/core/services"
	"task-management-user-service/pkg"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type server struct {
	App *fiber.App
	cfg *configs.Config
	db  *gorm.DB
}

func NewServer(config *configs.Config, db pkg.Database) *server {
	fiberCfg := fiber.Config{
		AppName: config.App.Name,
		// EnableTrustedProxyCheck: config.App.EnableTrustedProxyCheck,
		ProxyHeader: "X-Forwarded-For",
		// Prefork:                 true,
		ReadBufferSize:  1024 * 1024 * 15,
		WriteBufferSize: 1024 * 1024 * 15,
	}

	return &server{
		App: fiber.New(fiberCfg),
		cfg: config,
		db:  db.GetDB(),
	}
}

func (s *server) Start() {
	s.App.Use(s.Cors(s.cfg))

	api := s.App.Group("/api")

	userRepo := userrepo.NewUserRepository(s.db)
	userSrv := usersrv.NewUserService(userRepo)
	userCtl := userctl.NewUserController(userSrv)

	s.registerRoutes(userCtl, api)

	s.App.Get("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "ไม่พบเส้นทางที่ร้องขอ",
			"path":    c.Path(),
		})
	})

	fiberConnURL := fmt.Sprintf(":%s", s.cfg.App.Port)
	if strings.ToLower(s.cfg.App.Env) == "development" {
		fiberConnURL = "0.0.0.0:" + s.cfg.App.Port
	}
	if err := s.App.Listen(fiberConnURL); err != nil {
		panic(err.Error())
	}
}

func (s *server) Cors(cfg *configs.Config) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		corsConfig := cors.Config{
			AllowOrigins: cfg.App.Cors.AllowOrigins,
			AllowHeaders: "Content-Type, Authorization",
			AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
		}
		return cors.New(corsConfig)(c)
	}
}
