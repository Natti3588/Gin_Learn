package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ... トークン、 セッションなどをチェック
		c.Next()
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": "healthy!"})
}

func getProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"profile": "ないよ!"})
}

func updateSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"setting": "変更できないよ!"})
}

func main() {
	router := gin.Default()

	// 公開ルート
	public := router.Group("/api")
	{
		public.GET("/health", healthCheck)
	}

	// プライベートルート
	private := router.Group("/api")
	private.Use(AuthRequired())
	{
		private.GET("/profile", getProfile)
		private.GET("/settings", updateSettings)
	}

	router.Run(":8080")
}
