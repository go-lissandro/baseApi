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

func (cr CustommersRepository) CustommerById(cId string) models.Custommer {
	custummers := cr.Repository.GetByID(cId)

	return custummers
}

func (cr CustommersRepository) CustommerDelete(cId string) {
	cr.Repository.DeleteCustommer(cId)
}

func (cr CustommersRepository) UpdateCustommer(body models.Custommer, cId string) models.Custommer {
	custummer := cr.Repository.UpdateCustommer(body, cId)

	return custummer
}

func NewCustommersSevice(repository repository.CustommersRepository) *CustommersRepository {
	return &CustommersRepository{repository}
}
