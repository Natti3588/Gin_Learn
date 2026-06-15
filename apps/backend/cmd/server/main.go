package main

import (
	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/router"
	"backend/internal/service"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// @title           サンプル API
// @version         1.0
// @host            localhost:8080
func main() {
	_ = godotenv.Load()
	db := config.InitDB()
	if err := db.AutoMigrate(&model.Post{}); err != nil {
		panic(fmt.Errorf("マイグレーション失敗: %w", err))
	}

	// 依存性の組み立て
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	r := router.New(postHandler)

	// 起動用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		panic(fmt.Errorf("起動失敗: %w", err))
	}
}
