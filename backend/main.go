package main

import (
	"log"
	"net/http"
	"os"
	"ptihsan/config"
	"ptihsan/database"
	"ptihsan/logs"
	"ptihsan/router"
)

func main() {
	// Load konfigurasi dari .env
	config.LoadEnv()

	// Koneksi database
	if err := database.ConnectDB(); err != nil {
		logs.LogErrorSimple("Gagal koneksi ke database: " + err.Error())
		log.Fatal("Gagal koneksi ke database")
	}
	logs.LogInfoSimple("Koneksi ke database berhasil")

	// Setup route
	router.SetupRoutes()

	// Jalankan HTTP server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	logs.LogInfoSimple("Server berjalan di http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
