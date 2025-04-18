package config

import (
	"log"
	"uooobarry/soup/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	MigrateAll(db)
	log.Print("db is ready")

	return db
}

func MigrateAll(db *gorm.DB) error {
	models := []any{
		&model.User{},
		&model.Soup{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}
	return nil
}
