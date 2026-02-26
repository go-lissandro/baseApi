package models

type User struct {
	Base
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"uniqueIndex;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;unique"`
}
