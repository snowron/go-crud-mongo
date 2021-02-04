package interfaces

import "addressproject/model"

type IAddressService interface {
	GetAdressWithCoordinates(location model.Location) model.Address
	SearchFullTextAddress(keywords string) []model.Address
	AddAddress(address model.Address) string
	DeleteAddress(Id string)
	UpdateAddress(Id string, address model.Address)
}
