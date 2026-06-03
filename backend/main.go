package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids") // URLクエリ文字列をマップとして受け取る
		names := c.PostForm("names")

		fmt.Printf("ids: %v; names: %v\n", ids, names)
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	router.Run(":8080")
}
