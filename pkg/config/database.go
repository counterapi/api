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
		DSN:                  getDBDNS(),
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

// getDBDNS generates the database DNS.
func getDBDNS() string {
	dns := fmt.Sprintf(
		"host=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	if val, present := os.LookupEnv("DB_NAME"); present {
		dns += " dbname=" + val
	}

	if val, present := os.LookupEnv("DB_USER"); present {
		dns += " user=" + val
	}

	if val, present := os.LookupEnv("DB_PASSWORD"); present {
		dns += " password=" + val
	}

	return dns
}
