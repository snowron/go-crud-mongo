package controller

import (
	"addressproject/interfaces"
	"addressproject/model"
	"fmt"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddressController struct {
	AddressService interfaces.IAddressService
}

func (ee AddressController) SearchFullTextAddress(c *fiber.Ctx) error {
	return c.Status(200).JSON(ee.AddressService.SearchFullTextAddress(c.Query("keywords")))
}
func (ee AddressController) GetAdressWithCoordinates(c *fiber.Ctx) error {
	latitudeFloat, _ := strconv.ParseFloat(c.Query("Latitude"), 64)
	longitudeFloat, _ := strconv.ParseFloat(c.Query("Longitude"), 64)
	coor := model.Location{latitudeFloat, longitudeFloat}

	return c.Status(200).JSON(ee.AddressService.GetAdressWithCoordinates(coor))
}
func (ee AddressController) AddAddress(c *fiber.Ctx) error {
	address := model.Address{}
	if err := c.BodyParser(&address); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	fmt.Println(address)
	return c.Status(201).SendString(ee.AddressService.AddAddress(address))
}

func (ee AddressController) UpdateAddress(c *fiber.Ctx) error {
	address := model.Address{}
	idParam := c.Params("id")
	if err := c.BodyParser(&address); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	ee.AddressService.UpdateAddress(idParam, address)
	return c.SendStatus(204)
}

func (ee AddressController) DeleteAddress(c *fiber.Ctx) error {
	idParam := c.Params("id")
	ee.AddressService.DeleteAddress(idParam)
	return c.SendStatus(204)
}
func (ee AddressController) GetAddressWithId(c *fiber.Ctx) error {
	idParam := c.Query("id")
	ee.AddressService.DeleteAddress(idParam)
	return c.SendStatus(204)
}
