package server

import (
	"addressproject/config"
	"addressproject/controller"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	Config config.Configuration
}

func New(config config.Configuration) Server {
	server := Server{fiber.New(), config}
	return server
}
func (s Server) LoadRouter(controller controller.AddressController) Server {
	SetRouter(s.app, controller)
	return s
}

func (s Server) Run() {
	err := s.app.Listen(s.Config.Server.Port)
	if err != nil {
		println("wow")
	}

}
