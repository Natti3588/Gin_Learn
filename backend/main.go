package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		// c.PostForm() はリクエストから値を取得する
		message := c.PostForm("message")

		// c.DefaultPostForm() は値が存在しなかったら、第二引数の値をセットする
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run(":8080")
}
