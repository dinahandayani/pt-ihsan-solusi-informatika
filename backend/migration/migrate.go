package migration

import (
	"log"
	"ptihsan/database"
	"ptihsan/model"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration success")
}
