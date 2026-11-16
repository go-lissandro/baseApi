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
	if user.Email == "" || user.Username == "" || user.Password == "" {
		return errors.New("username, email e password sao obrigatorios")
	}

	// Antes usava DB.First(), que pegava um usuário qualquer e falhava com
	// "record not found" quando a tabela estava vazia (primeiro cadastro).
	var existing []models.User
	if err := udb.DB.Where("email = ? OR username = ?", user.Email, user.Username).Find(&existing).Error; err != nil {
		log.Printf("user userRepository: %s", err)
		return err
	}
	for _, u := range existing {
		if isUser, messageErr := existUser(u, user); isUser {
			return errors.New(messageErr)
		}
	}

	hash, err := handler.GenerateHashPassword(user.Password)
	if err != nil {
		log.Printf("erro hash %s", err)
		return err
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
