package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model

	ID        string     `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type User struct {
	Base
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
}

type Custommer struct {
	Base
	Body     string `gorm:"size:255;not null" json:"body"`
	TypeFile string `gorm:"size:255" json:"type_file"`
	File     string `gorm:"size:255" json:"file"`
}
