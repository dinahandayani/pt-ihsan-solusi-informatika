package main

import (
	"log"
	"ptihsan/config"
	"ptihsan/database"
	"ptihsan/migration"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()

	log.Println("Menjalankan migrasi...")
	migration.RunMigration()
}
