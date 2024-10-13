package repository

import "boiler-plate/types"

type UserRepository struct {
	UserMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
