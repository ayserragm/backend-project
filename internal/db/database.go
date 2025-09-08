package db

import (
	"fmt"
	"log"

	"github.com/ayserragm/backend-project/internal/config"
	"github.com/ayserragm/backend-project/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database: ", err)
	}

	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("❌ Migration error: ", err)
	}

	DB = database
	log.Println("✅ Database connected & migrated")
}
