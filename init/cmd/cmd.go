package cmd

import (
	"boiler-plate/config"
	"boiler-plate/network"
	"fmt"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	c.network.ServerStart(c.config.Server.Port)

	fmt.Println(c.config.Server.Port)

	return c
}
