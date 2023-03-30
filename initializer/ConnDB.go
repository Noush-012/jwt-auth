package initializer

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// To connect PSQL database
func ConnToDB() {

	dsn := os.Getenv("DATABASE")
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err == nil {
		log.Println("Database connected...!")
	} else {
		log.Println("Database connection failed...!")
	}

}
