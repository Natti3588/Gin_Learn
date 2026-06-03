package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// c.Query() / c.DefaultQuery() は URL クエリ文字列から読み取る
// c.PostForm() / c.DefaultPostForm() は application/x-www-form-urlencodedのリクエストボディから読み取る

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultPostForm("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
	})

	router.Run(":8080")
}
