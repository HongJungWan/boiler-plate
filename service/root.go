package service

import (
	"google.golang.org/protobuf/compiler/protogen"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *protogen.Service
)

type Service struct {
	// repository
}

func NewService() *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{}
	})

	return serviceInstance
}
