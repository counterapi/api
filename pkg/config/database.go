package config

import (
	"fmt"
	"os"

	"github.com/counterapi/counterapi/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global database variable.
var DB *gorm.DB //nolint:gochecknoglobals // allow global DB.

// SetupDatabase sets the database up.
func SetupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s TimeZone=Asia/Shanghai",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
		),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, err
	}

	err = db.AutoMigrate(&models.Counter{})
	if err != nil {
		return &gorm.DB{}, err
	}

	err = db.AutoMigrate(&models.Count{})
	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}
