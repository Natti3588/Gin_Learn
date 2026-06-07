package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()

	router.Any("/testing", startPage)
	router.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBindQuery(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Address,
	})
}
