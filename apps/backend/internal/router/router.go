package router

import (
	"backend/internal/handler"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New は gin.Engine を生成し、ルートと swagger を登録して返す。
// 依存（handler）は外から注入する。
func New(postHandler *handler.PostHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 投稿系（ブログ）
	posts := r.Group("/posts")
	{
		posts.POST("/", postHandler.Create)
		posts.GET("/", postHandler.List)
	}

	return r
}
