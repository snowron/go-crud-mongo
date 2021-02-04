package interfaces

import "addressproject/model"

type IAddressRepository interface {
	FindAddressWithCoordinates(model.Location) model.Address
	SearchAddressWithText(keywords string) []model.Address
	AddAddress(address model.Address) string
	DeleteAddress(Id string)
	UpdateAddress(Id string, address model.Address)
}
