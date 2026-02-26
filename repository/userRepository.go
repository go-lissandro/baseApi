package repository

import (
	"errors"
	"goRest/handler"
	"goRest/models"
	"log"

	"gorm.io/gorm"
)

type UsersDB struct {
	DB *gorm.DB
}

type UserRepository interface {
	UserInsert(custummer models.User) error
	FindUserByEmail(email string) (models.User, error)
}

func (udb UsersDB) UserInsert(user models.User) error {
	var useeers models.User
	pass := user.Password
	hash, err := handler.GenerateHashPassword(pass)

	if err != nil {
		log.Printf("erro hash %s", err)
		return err
	}
	resultUser := udb.DB.First(&useeers)
	if resultUser.Error != nil {
		log.Printf("user userRepository: %s", resultUser.Error)
		return resultUser.Error
	}
	isUser, messageErr := existUser(useeers, user)
	if isUser == true {
		return errors.New(messageErr)
	}

	user.Password = hash
	if result := udb.DB.Create(&user); result.Error != nil {
		log.Printf("user userRepository: %s", result.Error)

		return result.Error
	}

	return nil

}

func (udb UsersDB) FindUserByEmail(email string) (models.User, error) {
	var user models.User

	result := udb.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func existUser(user models.User, body models.User) (bool, string) {
	if user.Email == body.Email {
		return true, "E-mail existente"
	}

	if user.Username == body.Username {
		return true, "Usuario existente"
	}
	return false, ""
}
func NewUsersDB(dbClient *gorm.DB) *UsersDB {
	return &UsersDB{dbClient}
}
