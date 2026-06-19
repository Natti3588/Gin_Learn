package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Filters struct {
	Tags   []string `form:"tags" collection_format:"csv"`
	Labels []string `form:"labels" collection_format:"multi"`
	IdsSSV []int    `form:"ids_ssv" collection_format:"ssv"`
	IdsTSV []int    `form:"ids_tsv" collection_format:"tsv"`
	Levels []int    `form:"levels" collection_format:"pipes"`
}

func main() {
	router := gin.Default()

	router.GET("/search", func(c *gin.Context) {
		var f Filters
		if err := c.ShouldBind(&f); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, f)
	})
	router.Run(":8080")
}
