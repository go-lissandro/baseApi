package repository

import (
	"goRest/models"
	"log"

	"gorm.io/gorm"
)

type CustommersDB struct {
	DB *gorm.DB
}

type CustommersRepository interface {
	Insert(custommer models.Custommer) error
	GetAll() []models.Custommer
}

func (cdb CustommersDB) Insert(custommer models.Custommer) error {
	if result := cdb.DB.Create(&custommer); result.Error != nil {
		log.Printf("create custommerRepository: %s", result.Error)

		return result.Error
	}

	return nil
}

func (cdb CustommersDB) GetAll() []models.Custommer {

	var custommers []models.Custommer
	cdb.DB.Find(&custommers)

	return custommers
}

func (cdb CustommersDB) GetByID(cdbId string) models.Custommer {
	var custommer models.Custommer
	cdb.DB.Find(&custommer, cdbId)

	return custommer
}

func (cdb CustommersDB) DeleteCustommer(cbdId string) {
	var custommer models.Custommer
	cdb.DB.Delete(&custommer, cbdId)
}

func (cdb CustommersDB) UpdateCustommer(cdbId string, body models.Custommer) models.Custommer {
	var custommer models.Custommer

	cdb.DB.Find(&custommer, cdbId)
	cdb.DB.Model(&custommer)

	return custommer
}

func NewCustommersDB(dbClient *gorm.DB) *CustommersDB {
	return &CustommersDB{dbClient}
}
