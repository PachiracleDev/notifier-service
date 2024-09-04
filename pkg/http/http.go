package http

import (
	"fmt"
	"notifier/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HttpServer struct {
	app *fiber.App
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewHttpServer() *HttpServer {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
		BodyLimit: 1024 * 1024 * 20,
	})

	app.Use(cors.New())

	return &HttpServer{
		app: app,
	}
}

func RunHttpServer(server *HttpServer, conf *config.Config) error {

	server.app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			conf.Auth.Username: conf.Auth.Password,
		},
	}))

	go server.app.Listen(fmt.Sprintf(":%d", conf.Server.Port))
	return nil
}
func (u *HttpServer) App() *fiber.App {
	return u.app
}

func (u *HttpServer) BasicAuthMiddleware() {

}
