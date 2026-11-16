package models

type User struct {
	Base
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"uniqueIndex;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null"`
	// Role não pode vir do JSON de cadastro; promova admins direto no banco.
	Role string `gorm:"size:10;not null;default:user" json:"-"`
}
