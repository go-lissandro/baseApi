package configs

import (
	"goRest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgreSQL() *gorm.DB {
	dsn := EnvConfigDBUrl()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrateErr := db.AutoMigrate(&models.Custommer{}, &models.User{})
	if migrateErr != nil {
		panic(migrateErr)
	}

	return db
}
