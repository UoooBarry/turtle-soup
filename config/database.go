package config

import (
	"log"
	"uooobarry/soup/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Migrate all models
	AutoMigrate(db, &model.User{})
	AutoMigrate(db, &model.Soup{})

	model.DB = db
	log.Print("db is ready")
}

func AutoMigrate(db *gorm.DB, model any) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

}
