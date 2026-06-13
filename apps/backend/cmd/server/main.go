package main

import (
	"backend/internal/config"
	"backend/internal/model"
	"backend/internal/router"
	"fmt"
	"log"
	"os"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	engine := gin.Default()

	router.SetupRoutes(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 起動用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := engine.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
