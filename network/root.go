package network

import (
	"boiler-plate/service"
	"github.com/gin-gonic/gin"
)

type Network struct {
	engin   *gin.Engine
	service *service.Service
}

func NewNetwork(svc *service.Service) *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r, svc.User)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
