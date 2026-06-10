package main

import (
	"backend/config"
	"backend/routes"
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
	router := gin.Default()

	routes.SetupROutes(router, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 起動用
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
