package services

import (
	"goRest/models"
	"goRest/repository"
)

type CustommersRepository struct {
	Repository repository.CustommersRepository
}

type CustommerService interface {
	CustommerInsert(custommer models.Custommer) error
	CustommerGet() []models.Custommer
}

func (cr CustommersRepository) CustommerInsert(custommer models.Custommer) error {
	err := cr.Repository.Insert(custommer)

	if err != nil {
		return err
	}

	return nil
}

func (cr CustommersRepository) CustommerGet() []models.Custommer {
	custummers := cr.Repository.GetAll()

	return custummers
}

func NewCustommersSevice(repository repository.CustommersRepository) *CustommersRepository {
	return &CustommersRepository{repository}
}
