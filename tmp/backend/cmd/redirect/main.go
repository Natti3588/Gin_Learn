package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/old", func(c *gin.Context) {
		// StatusMovedPermanently は 301Moved Permanently
		// このリソースは恒久的に移動したという意味
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	})

	router.POST("/submit", func(c *gin.Context) {
		// StatusFound は 302Found 一時的な移動を表します
		c.Redirect(http.StatusFound, "/result")
	})

	router.GET("/test", func(c *gin.Context) {
		// c.Request.URL.Path = "/final" でリクエストのパスを書き換える
		// router.HandleContext(c)を呼ぶと、書き換えた後の /final で再度実行をする
		c.Request.URL.Path = "/final"
		router.HandleContext(c)
	})

	router.GET("/final", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.GET("/result", func(c *gin.Context) {
		c.String(http.StatusOK, "Redirected here!")
	})

	router.Run(":8080")
}
