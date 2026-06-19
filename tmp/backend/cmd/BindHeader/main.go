package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		h := testHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})
	router.Run(":8080")
}
