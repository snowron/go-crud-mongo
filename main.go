package main

import (
	"addressproject/config"
	"addressproject/controller"
	"addressproject/helper"
	"addressproject/interfaces"
	"addressproject/repo"
	"addressproject/server"
	"addressproject/service"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	helperDb := helper.MongoDB{}.Connect()
	var addressRepository interfaces.IAddressRepository = repo.AddressRepository{helperDb}

	var serviceAddress interfaces.IAddressService = service.AddressService{addressRepository}
	var controllerAddress = controller.AddressController{serviceAddress}

	var configuration config.Configuration
	viper.SetConfigName("config") // config file name without extension
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	err = viper.Unmarshal(&configuration)
	fmt.Println(configuration)

	server.New(configuration).LoadRouter(controllerAddress).Run()

}
