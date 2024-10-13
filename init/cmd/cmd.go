package cmd

import (
	"boiler-plate/config"
	"boiler-plate/network"
	"boiler-plate/repository"
	"boiler-plate/service"
	"fmt"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)
	c.network.ServerStart(c.config.Server.Port)

	fmt.Println(c.config.Server.Port)

	return c
}
