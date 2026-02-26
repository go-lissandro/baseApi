package services

import (
	"goRest/models"
	"goRest/repository"
)

type UserRepository struct {
	Reposotory repository.UserRepository
}

type UserService interface {
	UserInsert(body models.User) error
}

func (ur UserRepository) UserInsert(body models.User) error {
	new := ur.Reposotory.UserInsert(body)
	return new
}

func NewUserSevice(repository repository.UserRepository) *UserRepository {
	return &UserRepository{repository}
}
