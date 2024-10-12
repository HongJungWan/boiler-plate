package repository

import (
	"github.com/cloudwego/iasm/expr"
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *expr.Repository
)

type Repository struct {
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{}
	})

	return repositoryInstance
}
