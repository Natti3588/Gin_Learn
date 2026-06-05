package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/events", func(c *gin.Context) {
		cursor := c.Query("cursor")
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

		if limit > 100 {
			limit = 100
		}

		_ = cursor

		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"data":        []gin.H{},
			"next_cursor": "",
		})
	})

	router.Run(":8080")
}
