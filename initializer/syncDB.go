package initializer

import (
	"log"

	"github.com/Noush-012/JWT/models"
)

func SyncDatabase() {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Println("DB sync failed..!")
	}
	log.Println("DB sync successful..!")
}
