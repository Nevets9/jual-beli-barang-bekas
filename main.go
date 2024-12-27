package main

import (
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/routes"
	"jual-beli-barang-bekas/models"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
)

func main() {
	// Inisialisasi database
	db := config.InitDB()

	// Auto-migrate models (untuk mengatur schema)
	err := db.AutoMigrate(&models.User{}, &models.Item{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Migrasi gagal:", err)
	}

	// Inisialisasi router Gin
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Menjalankan server
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server gagal dijalankan:", err)
	}

	fmt.Println("Server berjalan di http://localhost:8080")
}
