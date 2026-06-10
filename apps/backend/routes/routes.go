package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupROutes(router *gin.Engine) {
	// 投稿系（ブログ）
	posts := router.Group("/posts")
	{
		posts.POST("/chat", controllers.ChatHandler)
	}

}