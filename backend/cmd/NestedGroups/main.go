package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{"id": "1", "name": "Natti"},
		{"id": "2", "name": "Sakana"},
	})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"user": id})
}

func listPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"post": "success"})
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"getpost": id})
}

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		// /api/v1
		v1 := api.Group("/v1")
		{
			// /api/v1/users
			users := v1.Group("/users")
			users.GET("/", listUsers)
			users.GET("/:id", getUser)

			// /api/posts
			posts := v1.Group("/posts")
			posts.GET("/", listPosts)
			posts.GET("/:id", getPost)
		}
	}

	router.Run(":8080")
}
