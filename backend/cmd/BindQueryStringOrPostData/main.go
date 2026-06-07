package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"Birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	router := gin.Default()

	router.GET("/testing", startPage)
	router.POST("/testing", startPage)
	router.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Name: %s, Address: %s, Birthday: %s\n", person.Name, person.Address, person.Birthday)
	c.JSON(http.StatusOK, gin.H{
		"name":     person.Name,
		"address":  person.Address,
		"birthday": person.Birthday,
	})
}
