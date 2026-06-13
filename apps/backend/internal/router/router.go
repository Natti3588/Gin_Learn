package router

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// 投稿系（ブログ）
	posts := router.Group("/posts")
	{
		posts.POST("/chat", handler.ChatHandler)
	}
}
