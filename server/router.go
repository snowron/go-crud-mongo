package server

import (
	"addressproject/controller"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
}

func SetRouter(app *fiber.App, controllerAddress controller.AddressController) {
	app.Get("/v1/Search/fulltext", controllerAddress.SearchFullTextAddress)
	app.Get("/v1/Search/coordinates", controllerAddress.GetAdressWithCoordinates)

	app.Get("/v1/Address/:id", controllerAddress.GetAddressWithId)
	app.Delete("/v1/Address/:id", controllerAddress.DeleteAddress)
	app.Put("/v1/Address/:id", controllerAddress.UpdateAddress)
	app.Post("/v1/Address", controllerAddress.AddAddress)
}
