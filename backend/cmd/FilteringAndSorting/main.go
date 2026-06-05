package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/products", func(c *gin.Context) {
		category := c.Query("category")
		minPrice := c.Query("min_price")
		maxPrice := c.Query("max_price")
		sortBy := c.DefaultQuery("sort", "created_at")
		order := c.DefaultQuery("order", "desc")

		allowed := map[string]bool{"created_at": true, "price": true, "name": true}
		if !allowed[sortBy] {
			sortBy = "desc"
		}
		if order != "asc" && order != "desc" {
			order = "desc"
		}

		_ = category
		_ = minPrice
		_ = maxPrice

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    []gin.H{},
			"filters": gin.H{
				"category":  category,
				"min_price": minPrice,
				"max_price": maxPrice,
				"sort":      sortBy,
				"order":     order,
			},
		})
	})

	router.Run(":8080")
}
