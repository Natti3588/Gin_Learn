package main

import (
	"backend/internal/config"
	"backend/internal/model"
	"backend/internal/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title           サンプル API
// @version         1.0
// @host            localhost:8080
func main() {
	godotenv.Load()
	db := config.InitDB()
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		panic(fmt.Errorf("マイグレーション失敗: %w", err))
	}

	r := router.New()

	// 起動用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
