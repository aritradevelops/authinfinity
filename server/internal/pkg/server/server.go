package server

import (
	"context"
	"fmt"
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/app"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/auth"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/emailverificationrequest"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/oauth"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/password"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/resetpasswordrequest"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/session"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/middlewares/errorhandler"
	"github.com/aritradevelops/authinfinity/server/internal/middlewares/translator"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	cfg *config.Config
	db  db.Database
	app *fiber.App
}

func New(cfg *config.Config, db db.Database) *Server {

	app := fiber.New(fiber.Config{
		AppName:      cfg.Env.ServiceName,
		ErrorHandler: errorhandler.New(),
	})

	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(translator.New())

	server := &Server{
		cfg: cfg,
		db:  db,
		app: app,
	}

	server.setupRoutes()

	return server
}
func (s *Server) Start() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.cfg.Env.Port))
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func (s *Server) setupRoutes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello!",
			"name":    s.cfg.Env.ServiceName,
			"status":  "running",
		})
	})
	s.app.Get("/health", func(c *fiber.Ctx) error {
		err := s.db.Health()
		status := "healthy"
		if err != nil {
			status = "unhealthy"
		}
		return c.JSON(fiber.Map{
			"status":    status,
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	apiV1 := s.app.Group("/api/v1")
	resetpasswordrequest.

		// ------ DO NOT MODIFY THE BELOW Manually. THIS IS AUTO GENERATED ------ //
		RegisterRoutes(apiV1)
	user.RegisterRoutes(apiV1)
	auth.RegisterRoutes(apiV1)
	account.RegisterRoutes(apiV1)
	app.RegisterRoutes(apiV1)
	oauth.RegisterRoutes(apiV1)
	password.RegisterRoutes(apiV1)
	session.RegisterRoutes(apiV1)
	emailverificationrequest.RegisterRoutes(apiV1)

}
