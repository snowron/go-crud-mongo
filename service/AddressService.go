package service

import (
	"addressproject/interfaces"
	"addressproject/model"
)

type AddressService struct {
	Repository interfaces.IAddressRepository
}

func (s AddressService) GetAdressWithCoordinates(location model.Location) model.Address {

	return s.Repository.FindAddressWithCoordinates(location)
}
func (s AddressService) SearchFullTextAddress(keywords string) []model.Address {
	return s.Repository.SearchAddressWithText(keywords)

}

func (s AddressService) AddAddress(address model.Address) string {

	return s.Repository.AddAddress(address)
}

func (s AddressService) DeleteAddress(Id string) {
	s.Repository.DeleteAddress(Id)
}

func (s AddressService) UpdateAddress(Id string, address model.Address) {
	s.Repository.UpdateAddress(Id, address)
}
